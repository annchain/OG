package consensus

import (
	"fmt"
	"github.com/annchain/OG/arefactor/consensus_interface"
	"github.com/annchain/OG/arefactor/transport_interface"
	"github.com/latifrons/goffchan"
	"github.com/latifrons/soccerdash"
	"github.com/sirupsen/logrus"
	"strconv"
)

/**
Implemented according to
HotStuff: BFT Consensus in the Lens of Blockchain
Maofan Yin, Dahlia Malkhi, Michael K. Reiter, Guy Golan Gueta and Ittai Abraham
*/

type Partner struct {
	//MessageHub       Hub
	//Ledger           *Ledger
	PeerIds   []string
	MyIdIndex int
	N         int
	F         int
	PaceMaker *PaceMaker
	Safety    *Safety
	//BlockTree        *BlockTree
	Logger *logrus.Logger
	Report *soccerdash.Reporter
	quit   chan bool

	ProposalContextProvider consensus_interface.ProposalContextProvider
	ProposalGenerator       consensus_interface.ProposalGenerator
	ProposalVerifier        consensus_interface.ProposalVerifier
	ProposalExecutor        consensus_interface.ProposalExecutor
	CommitteeProvider       consensus_interface.CommitteeProvider
	Signer                  consensus_interface.Signer
	AccountProvider         consensus_interface.AccountHolder
	Hasher                  consensus_interface.Hasher
	Ledger                  consensus_interface.Ledger

	pendingBlockTree *PendingBlockTree
	pendingVotes     map[string]consensus_interface.SignatureCollector // collected votes per block indexed by their LedgerInfo hash

	// event handlers
	myNewIncomingMessageEventChan chan *transport_interface.IncomingLetter
	newOutgoingMessageSubscribers []transport_interface.NewOutgoingMessageEventSubscriber // a message need to be sent

}

func (n *Partner) InitDefault() {
	n.quit = make(chan bool)
	// for each hash, init a SignatureCollector
	n.pendingBlockTree = &PendingBlockTree{
		Logger: n.Logger,
		Ledger: n.Ledger,
	}
	n.pendingVotes = make(map[string]consensus_interface.SignatureCollector)
	n.myNewIncomingMessageEventChan = make(chan *transport_interface.IncomingLetter)
	n.newOutgoingMessageSubscribers = []transport_interface.NewOutgoingMessageEventSubscriber{}
}
func (n *Partner) Start() {
	for {
		logrus.Trace("partner loop round start")
		select {
		case <-n.quit:
			return
		case msg := <-n.myNewIncomingMessageEventChan:
			n.handleIncomingMessage(msg)

			//n.Logger.WithField("msgType", msg.Typev.HotStuffMessageTypeString()).WithField("msgc", msg).Info("received message")
			//if ok := n.signatureOk(msg); !ok {
			//	fmt.Println(msg)
			//	panic("signature invalid")
			//	//continue
			//}

		case <-n.PaceMaker.timer.C:
			logrus.WithField("round", n.PaceMaker.CurrentRound).Warn("timeout")
			n.PaceMaker.LocalTimeoutRound()
		}
		logrus.Trace("partner loop round end")
	}
}

func (n *Partner) Stop() {
	close(n.quit)
}

func (n *Partner) Name() string {
	return fmt.Sprintf("Node %d", n.MyIdIndex)
}

func (n *Partner) ProcessProposalMessage(msg *consensus_interface.HotStuffSignedMessage) {

	p := &consensus_interface.ContentProposal{}
	err := p.FromBytes(msg.ContentBytes)
	if err != nil {
		logrus.WithError(err).Debug("failed to decode ContentProposal")
		return
	}

	n.ProcessCertificates(p.Proposal.ParentQC, p.TC, "ProposalM")

	currentRound := n.PaceMaker.CurrentRound

	if p.Proposal.Round != currentRound {
		n.Logger.WithField("pRound", p.Proposal.Round).WithField("currentRound", currentRound).Warn("current round not match.")
		return
	}

	if msg.SenderId != n.CommitteeProvider.GetLeaderPeerId(currentRound) {
		n.Logger.WithField("msg.SenderId", msg.SenderId).
			WithField("current leader", n.CommitteeProvider.GetLeaderPeerId(currentRound)).
			Warn("current leader not match.")
		return
	}
	// verify proposal
	// TODO: now sync. change to async in the future
	verifyResult := n.ProposalVerifier.VerifyProposal(p)
	if !verifyResult.Ok {
		logrus.Debug("proposal verification failed")
		return
	}

	// execute the block
	// TODO: execute the block async
	//n.BlockTree.ExecuteAndInsert(&p.HotStuffMessageTypeProposal)
	// TODO: who is ProposalExecutor?
	n.ProposalExecutor.ExecuteProposal(&p.Proposal)

	// vote after execution

	voteMsg := n.Safety.MakeVote(p.Proposal.Id, p.Proposal.Round, p.Proposal.ParentQC)
	if voteMsg != nil {
		bytes := voteMsg.ToBytes()
		voteAggregator := n.CommitteeProvider.GetLeaderPeerId(currentRound + 1)

		signature, err := n.sign(voteMsg)
		if err != nil {
			return
		}

		outMsg := &consensus_interface.HotStuffSignedMessage{
			HotStuffMessageType: int(consensus_interface.HotStuffMessageTypeVote),
			ContentBytes:        bytes,
			SenderId:            n.CommitteeProvider.GetMyPeerId(),
			Signature:           signature,
		}
		letter := &transport_interface.OutgoingLetter{
			Msg:            outMsg,
			SendType:       transport_interface.SendTypeUnicast,
			CloseAfterSent: false,
			EndReceivers:   []string{voteAggregator},
		}

		n.notifyNewOutgoingMessage(letter)
	}
}

func (n *Partner) ProcessVoteMessage(msg *consensus_interface.HotStuffSignedMessage) {
	p := &consensus_interface.ContentVote{}
	err := p.FromBytes(msg.ContentBytes)
	if err != nil {
		logrus.WithError(err).Debug("failed to decode ContentVote")
		return
	}
	n.ProcessCertificates(p.QC, p.TC, "Vote")
	n.ProcessVote(p, msg.Signature, msg.SenderId)
}

func (t *Partner) ProcessVote(vote *consensus_interface.ContentVote, signature consensus_interface.Signature, fromId string) {
	id, err := t.CommitteeProvider.GetPeerIndex(fromId)
	if err != nil {
		logrus.WithError(err).WithField("peerId", fromId).
			Debug("error in finding peer in committee")
	}

	voteIndex := t.Hasher.Hash(vote.LedgerCommitInfo.GetHashContent())

	collector := t.ensureCollector(voteIndex)
	collector.Collect(signature, id)

	logrus.WithField("sigs", collector.GetCurrentCount()).
		WithField("sig", signature).Debug("signature got one")

	if collector.Collected() {
		t.Logger.WithField("vote", vote).Info("votes collected")
		qc := &consensus_interface.QC{
			VoteData:       vote.VoteInfo, // TODO: check if the voteinfo is aligned
			JointSignature: collector.GetJointSignature(),
		}

		t.pendingBlockTree.EnsureHighQC(qc)
		t.PaceMaker.AdvanceRound(qc, nil, "vote qc got")

	} else {
		t.Logger.WithField("vote", vote).
			WithField("now", collector.GetCurrentCount()).Trace("votes yet collected")
	}
}

func (n *Partner) ProcessCertificates(qc *consensus_interface.QC, tc *consensus_interface.TC, reason string) {
	n.PaceMaker.AdvanceRound(qc, tc, reason+"ProcessCertificates"+strconv.Itoa(n.PaceMaker.CurrentRound))
	if qc != nil {
		n.Safety.UpdatePreferredRound(qc)
		if qc.VoteData.ExecStateId != "" {
			n.pendingBlockTree.Commit(qc.VoteData.Id)
		}
	}
}

func (n *Partner) ProcessNewRoundEvent() {
	if n.CommitteeProvider.GetMyPeerId() != n.CommitteeProvider.GetLeaderPeerId(n.PaceMaker.CurrentRound) {
		// not the leader
		n.Logger.Trace("I'm not the leader so just return")
		return
	}
	//proposal := n.BlockTree.GenerateProposal(n.PaceMaker.CurrentRound, strconv.Itoa(RandInt()))
	proposalContext := n.ProposalContextProvider.GetProposalContext()

	proposal := n.ProposalGenerator.GenerateProposal(proposalContext)
	n.Logger.WithField("proposal", proposal).Warn("I'm the current leader")
	n.Report.Report("leader", proposal.Proposal.Round, false)

	bytes := proposal.ToBytes()
	signature, err := n.sign(proposal)
	if err != nil {
		return
	}

	// announce it
	outMsg := &consensus_interface.HotStuffSignedMessage{
		HotStuffMessageType: int(consensus_interface.HotStuffMessageTypeProposal),
		ContentBytes:        bytes,
		SenderId:            n.CommitteeProvider.GetMyPeerId(),
		Signature:           signature,
	}
	letter := &transport_interface.OutgoingLetter{
		Msg:            outMsg,
		SendType:       transport_interface.SendTypeUnicast,
		CloseAfterSent: false,
		EndReceivers:   n.CommitteeProvider.GetAllMemberPeedIds(),
	}
	n.notifyNewOutgoingMessage(letter)
}

func (n *Partner) SaveConsensusState() {
	n.Logger.WithFields(logrus.Fields{
		"lastVoteRound":  n.Safety.lastVoteRound,
		"preferredRound": n.Safety.preferredRound,
	}).Trace("Persist" + strconv.Itoa(n.PaceMaker.CurrentRound))
}

func (n *Partner) handleIncomingMessage(msg *transport_interface.IncomingLetter) {
	if msg.Msg.MsgType != consensus_interface.HotStuffMessageTypeRoot {
		return
	}
	// convert from wireMessage to SignedMessage since consensus need to verify signature
	signedMessage := &consensus_interface.HotStuffSignedMessage{}
	_, err := signedMessage.UnmarshalMsg(msg.Msg.ContentBytes)
	if err != nil {
		logrus.WithError(err).Debug("failed to parse HotStuffSignedMessage message")
		return
	}
	// TODO: verify if the sender is in the committee.
	// TODO: verify signature

	switch consensus_interface.HotStuffMessageType(signedMessage.HotStuffMessageType) {
	case consensus_interface.HotStuffMessageTypeProposal:
		logrus.Info("handling proposal")
		n.ProcessProposalMessage(signedMessage)
	case consensus_interface.HotStuffMessageTypeVote:
		logrus.Info("handling vote")
		n.ProcessVoteMessage(signedMessage)
	case consensus_interface.HotStuffMessageTypeTimeout:
		logrus.Info("handling timeout")
		n.PaceMaker.ProcessRemoteTimeout(signedMessage)
	default:
		panic("unsupported typev")
	}
}

// notifications

func (d *Partner) NewIncomingMessageEventChannel() chan *transport_interface.IncomingLetter {
	return d.myNewIncomingMessageEventChan
}

// subscribe mine
func (d *Partner) AddSubscriberNewOutgoingMessageEvent(sub transport_interface.NewOutgoingMessageEventSubscriber) {
	d.newOutgoingMessageSubscribers = append(d.newOutgoingMessageSubscribers, sub)
}

func (d *Partner) notifyNewOutgoingMessage(event *transport_interface.OutgoingLetter) {
	for _, subscriber := range d.newOutgoingMessageSubscribers {
		<-goffchan.NewTimeoutSenderShort(subscriber.NewOutgoingMessageEventChannel(), event, "outgoing hotstuff partner"+subscriber.Name()).C
		//subscriber.NewOutgoingMessageEventChannel() <- event
	}
}

func (n *Partner) ensureCollector(commitInfoHash string) consensus_interface.SignatureCollector {
	if _, ok := n.pendingVotes[commitInfoHash]; !ok {
		collector := &BlsSignatureCollector{
			Threshold: 2*n.F + 1,
		}
		collector.InitDefault()
		n.pendingVotes[commitInfoHash] = collector
	}
	collector := n.pendingVotes[commitInfoHash]
	return collector
}

func (n *Partner) sign(msg Signable) (signature []byte, err error) {
	privateKey, err := n.AccountProvider.ProvidePrivateKey(false)
	if err != nil {
		logrus.WithError(err).Warn("account provider cannot provide private key")
		return
	}
	signature = n.Signer.Sign(msg.SignatureTarget(), privateKey)
	return
}