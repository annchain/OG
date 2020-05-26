package community

import (
	"context"
	"errors"
	"fmt"
	"github.com/annchain/OG/ffchan"
	"github.com/libp2p/go-libp2p"
	core "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	swarm "github.com/libp2p/go-libp2p-swarm"
	"github.com/multiformats/go-multiaddr"
	"github.com/sirupsen/logrus"
	"github.com/tinylib/msgp/msgp"
	"log"
	"math/rand"
	"sync"
	"time"
)

var BackoffConnect = time.Second * 5
var IpLayerProtocol = "tcp"

type IoEvent struct {
	Neighbour *Neighbour
	Err       error
}

type Neighbour struct {
	Id              peer.ID
	Stream          network.Stream
	IoEventChannel  chan *IoEvent
	IncomingChannel chan *WireMessage
	msgpReader      *msgp.Reader
	msgpWriter      *msgp.Writer
	event           chan bool
	outgoingChannel chan *Msg
	quit            chan bool
}

func (c *Neighbour) InitDefault() {
	c.event = make(chan bool)
	c.quit = make(chan bool)
	c.outgoingChannel = make(chan *Msg) // messages already dispatched
}

func (c *Neighbour) StartRead() {
	var err error
	c.msgpReader = msgp.NewReader(c.Stream)
	for {
		msg := &WireMessage{}
		err = msg.DecodeMsg(c.msgpReader)
		if err != nil {
			// bad message, drop
			logrus.WithError(err).Warn("read error")
			break
		}

		<-ffchan.NewTimeoutSenderShort(c.IncomingChannel, msg, "read").C
		//c.IncomingChannel <- msg
	}
	// neighbour disconnected, notify the communicator
	c.IoEventChannel <- &IoEvent{
		Neighbour: c,
		Err:       err,
	}

}

func (c *Neighbour) StartWrite() {
	var err error
	c.msgpWriter = msgp.NewWriter(c.Stream)
loop:
	for {
		select {
		case req := <-c.outgoingChannel:
			logrus.Trace("neighbour got send request")
			contentBytes, err := req.Content.MarshalMsg([]byte{})
			if err != nil {
				panic(err)
			}

			wireMessage := WireMessage{
				MsgType:      int(req.Typev),
				ContentBytes: contentBytes,
				SenderId:     req.SenderId,
				Signature:    req.Sig,
			}

			err = wireMessage.EncodeMsg(c.msgpWriter)
			if err != nil {
				break
			}
			err = c.msgpWriter.Flush()
			if err != nil {
				break
			}
			logrus.Trace("neighbour sent")

		case <-c.quit:
			break loop
		}
	}
	// neighbour disconnected, notify the communicator
	c.IoEventChannel <- &IoEvent{
		Neighbour: c,
		Err:       err,
	}
}

func (c *Neighbour) Send(req *Msg) {
	<-ffchan.NewTimeoutSenderShort(c.outgoingChannel, req, "send").C
	//c.outgoingChannel <- req
}

// LibP2pPhysicalCommunicator
type LibP2pPhysicalCommunicator struct {
	Port       int // listening port
	PrivateKey core.PrivKey

	node            host.Host              // p2p host to receive new streams
	activePeers     map[peer.ID]*Neighbour // active peers that will be reconnect if error
	tryingPeers     map[peer.ID]bool       // peers that is trying to connect.
	incomingChannel chan *WireMessage      // incoming message channel
	outgoingChannel chan *OutgoingRequest  // universal outgoing channel to collect send requests
	ioEventChannel  chan *IoEvent          // receive event when peer disconnects

	initWait sync.WaitGroup
	quit     chan bool
	mu       sync.RWMutex
}

func (c *LibP2pPhysicalCommunicator) InitDefault() {
	c.activePeers = make(map[peer.ID]*Neighbour)
	c.tryingPeers = make(map[peer.ID]bool)
	c.outgoingChannel = make(chan *OutgoingRequest)
	c.incomingChannel = make(chan *WireMessage)
	c.ioEventChannel = make(chan *IoEvent)
	c.initWait.Add(1)
	c.quit = make(chan bool)
}

func (c *LibP2pPhysicalCommunicator) Start() {
	// start consuming queue
	go c.Listen()
	go c.consumeQueue()
	go c.recoverPeer()
}

func (c *LibP2pPhysicalCommunicator) consumeQueue() {
	c.initWait.Wait()
	for {
		select {
		case req := <-c.outgoingChannel:
			logrus.WithField("req", req).Trace("physical communicator got a request from outgoing channel")
			go c.handleRequest(req)
		case <-c.quit:
			return
		}
	}

}

func (c *LibP2pPhysicalCommunicator) Stop() {
	close(c.quit)
	// shut the node down
	if err := c.node.Close(); err != nil {
		panic(err)
	}
}

func (c *LibP2pPhysicalCommunicator) GetIncomingChannel() chan *WireMessage {
	return c.incomingChannel
}

func (c *LibP2pPhysicalCommunicator) makeHost(priv core.PrivKey) (host.Host, error) {
	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/%s/%d", IpLayerProtocol, c.Port)),
		libp2p.Identity(priv),
		libp2p.DisableRelay(),
	}
	basicHost, err := libp2p.New(context.Background(), opts...)
	if err != nil {
		return nil, err
	}
	return basicHost, nil
}

func (c *LibP2pPhysicalCommunicator) printHostInfo(basicHost host.Host) {

	// print the node's listening addresses
	// protocol is always p2p
	hostAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/p2p/%s", basicHost.ID().Pretty()))
	if err != nil {
		panic(err)
	}

	addr := basicHost.Addrs()[0]
	fullAddr := addr.Encapsulate(hostAddr)
	log.Printf("I am %s\n", fullAddr)
}

func (c *LibP2pPhysicalCommunicator) Listen() {
	// start a libp2p node with default settings
	host, err := c.makeHost(c.PrivateKey)
	if err != nil {
		panic(err)
	}
	c.node = host

	c.printHostInfo(c.node)

	c.node.SetStreamHandler(ProtocolId, c.HandlePeerStream)
	c.initWait.Done()
	logrus.Info("waiting for connection...")
	select {
	case <-c.quit:
		err := c.node.Close()
		if err != nil {
			logrus.WithError(err).Warn("closing communicator")
		}
	}
}

func (c *LibP2pPhysicalCommunicator) HandlePeerStream(s network.Stream) {
	// must lock to prevent double channel
	c.mu.Lock()
	defer c.mu.Unlock()

	logrus.WithFields(logrus.Fields{
		"peerId":  s.Conn().RemotePeer().String(),
		"address": s.Conn().RemoteMultiaddr().String(),
	}).Info("Peer connection established")
	peerId := s.Conn().RemotePeer()

	// deregister in the trying list
	delete(c.tryingPeers, s.Conn().RemotePeer())

	// prevent double channel
	if _, ok := c.activePeers[peerId]; ok {
		// already established. close
		err := s.Close()
		if err != nil {
			logrus.WithError(err).Warn("closing peer")
		}
		return
	}

	neightbour := &Neighbour{
		Id:              peerId,
		Stream:          s,
		IoEventChannel:  c.ioEventChannel,
		IncomingChannel: c.incomingChannel,
	}
	neightbour.InitDefault()
	c.activePeers[peerId] = neightbour

	go neightbour.StartRead()
	go neightbour.StartWrite()
}

func (c *LibP2pPhysicalCommunicator) ClosePeer(id string) {

}

func (c *LibP2pPhysicalCommunicator) GetNeighbour(id string) (neighbour *Neighbour, err error) {
	idp, err := peer.Decode(id)
	if err != nil {
		return
	}
	neighbour, ok := c.activePeers[idp]
	if !ok {
		err = errors.New("peer not active")
	}
	return
}

// SuggestConnection takes a peer address and try to connect to it.
func (c *LibP2pPhysicalCommunicator) SuggestConnection(address string) (peerIds string) {
	c.initWait.Wait()
	logrus.WithField("address", address).Info("registering address")
	fullAddr, err := multiaddr.NewMultiaddr(address)
	if err != nil {
		logrus.WithField("address", address).WithError(err).Warn("bad address")
		return
	}

	// p2p layer address
	p2pAddr, err := fullAddr.ValueForProtocol(multiaddr.P_P2P)

	if err != nil {
		logrus.WithField("address", address).WithError(err).Warn("bad address")
	}

	protocolAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/p2p/%s", p2pAddr))

	if err != nil {
		logrus.WithField("address", address).WithError(err).Warn("bad address")
	}
	// keep only the connection info, wipe out the p2p layer
	connectionAddr := fullAddr.Decapsulate(protocolAddr)
	//fmt.Println("connectionAddr:" + connectionAddr.String())

	// recover peerId from Base58 Encoded p2pAddr
	peerId, err := peer.Decode(p2pAddr)
	if err != nil {
		logrus.WithField("address", address).WithError(err).Warn("bad address")
	}
	peerIds = peerId.String()

	//fmt.Println("peerId:" + p2pAddr)
	// check if it is a self connection.
	if peerId == c.node.ID() {
		return
	}

	// save address and peer info
	c.node.Peerstore().AddAddr(peerId, connectionAddr, peerstore.PermanentAddrTTL)

	// reg in the trying list
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.tryingPeers[peerId]; ok {
		return
	}
	c.tryingPeers[peerId] = true
	return
}

func (c *LibP2pPhysicalCommunicator) Enqueue(req *OutgoingRequest) {
	logrus.WithField("req", req).Info("Sending message")
	c.initWait.Wait()
	<-ffchan.NewTimeoutSenderShort(c.outgoingChannel, req, "enqueue").C
	//c.outgoingChannel <- req
}

// we use direct connection currently so let's build a connection if not exists.
func (c *LibP2pPhysicalCommunicator) handleRequest(req *OutgoingRequest) {
	logrus.Trace("handling send request")
	if req.SendType == SendTypeBroadcast {
		for _, neighbour := range c.activePeers {
			go neighbour.Send(req.Msg)
		}
		return
	}
	// find neighbour first
	for _, peerIdEncoded := range req.EndReceivers {
		peerId, err := peer.Decode(peerIdEncoded)
		if err != nil {
			logrus.WithError(err).WithField("peerIdEncoded", peerIdEncoded).Warn("decoding peer")
		}
		logrus.WithField("peerId", peerId).Trace("handling sending requets")
		// get active neighbour
		neighbour, ok := c.activePeers[peerId]
		if !ok {
			// wait for node to be connected. currently node address are pre-located and connections are built ahead.
			logrus.WithField("peerId", peerId).Trace("connection not in active peers")
			continue
		}
		logrus.WithField("peerId", peerId).Trace("go send")
		go neighbour.Send(req.Msg)
	}
}

func (c *LibP2pPhysicalCommunicator) pickOneAndConnect() {
	c.mu.RLock()

	if len(c.tryingPeers) == 0 {
		c.mu.RUnlock()
		return
	}
	peerIds := []peer.ID{}
	for k, _ := range c.tryingPeers {
		peerIds = append(peerIds, k)
	}
	c.mu.RUnlock()

	peerId := peerIds[rand.Intn(len(peerIds))]

	// start a stream
	logrus.WithField("peerId", peerId).Trace("connecting peer")
	s, err := c.node.NewStream(context.Background(), peerId, ProtocolId)
	if err != nil {
		if err != swarm.ErrDialBackoff {
			//logrus.WithField("stream", s).WithError(err).Warn("error on starting stream")
		}
		return
	}
	// stream built
	// release the lock
	c.HandlePeerStream(s)

	//hub.handleStream(s)
	//// Create a buffered stream so that read and writes are non blocking.
	//rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
	//
	//// Create a thread to read and write data.
	//go hub.writeData(rw)
	//go hub.readData(rw)
	//logrus.WithField("s", info).WithError(err).Warn("connection established")
}

func (c *LibP2pPhysicalCommunicator) recoverPeer() {
	for {
		select {
		case <-c.quit:
			return
		case event := <-c.ioEventChannel:
			c.reportPeerDown(event.Neighbour)
		default:
			time.Sleep(time.Second * 1) // at least sleep one second to prevent flood
			c.pickOneAndConnect()
		}
	}
}

func (c *LibP2pPhysicalCommunicator) reportPeerDown(neighbour *Neighbour) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.activePeers, neighbour.Id)
	c.tryingPeers[neighbour.Id] = true
}
