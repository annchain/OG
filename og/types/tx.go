package types

import (
	"fmt"
	"github.com/annchain/OG/common"
	"github.com/annchain/OG/common/crypto"
	"github.com/annchain/OG/common/hexutil"
	"github.com/annchain/OG/common/math"
	"github.com/annchain/OG/types"
	"golang.org/x/crypto/sha3"
)

type Tx struct {
	// graph structure info
	Hash        common.Hash
	ParentsHash common.Hashes
	Height      uint64
	MineNonce   uint64
	Weight      uint64

	// tx info
	AccountNonce uint64
	From         *common.Address
	To           common.Address
	Value        *math.BigInt
	TokenId      int32
	Data         []byte
	PublicKey    crypto.PublicKey
	Signature    hexutil.Bytes
	//confirm     time.Time
	// CanRecoverPubFromSig
}

func (t Tx) CalcTxHash() (hash common.Hash) {
	w := types.NewBinaryWriter()

	for _, ancestor := range t.ParentsHash {
		w.Write(ancestor.Bytes)
	}
	// do not use Height to calculate tx hash.
	w.Write(t.Weight)
	w.Write(t.CalcMinedHash().Bytes)

	result := sha3.Sum256(w.Bytes())
	hash.MustSetBytes(result[0:], common.PaddingNone)
	return
}

func (t Tx) CalcMinedHash() (hash common.Hash) {
	w := types.NewBinaryWriter()
	//if !CanRecoverPubFromSig {
	w.Write(t.PublicKey)
	//}
	w.Write(t.Signature, t.MineNonce)
	result := sha3.Sum256(w.Bytes())
	hash.MustSetBytes(result[0:], common.PaddingNone)
	return
}

func (t *Tx) SignatureTargets() []byte {
	// log.WithField("tx", t).Tracef("SignatureTargets: %s", t.Dump())

	w := types.NewBinaryWriter()

	w.Write(t.AccountNonce)
	w.Write(t.From.Bytes)
	w.Write(t.To.Bytes, t.Value.GetSigBytes(), t.Data, t.TokenId)
	return w.Bytes()
}

func (t Tx) GetType() TxBaseType {
	return TxBaseTypeNormal
}

func (t Tx) GetHeight() uint64 {
	return t.Height
}

func (t Tx) GetWeight() uint64 {
	if t.Weight == 0 {
		panic("implementation error: weight not initialized")
	}
	return t.Weight
}

func (t Tx) GetTxHash() common.Hash {
	if t.Hash.Empty() {
		t.CalcTxHash()
	}
	return t.Hash
}

func (t Tx) Parents() common.Hashes {
	return t.ParentsHash
}

func (t Tx) String() string {
	return fmt.Sprintf("T-%d-[%.10s]-%d", t.Height, t.GetTxHash().Hex(), t.Weight)
}

//CalculateWeight  a core algorithm for tx sorting,
//a tx's weight must bigger than any of it's parent's weight  and bigger than any of it's elder transaction's
func (t Tx) CalculateWeight(parents Txis) uint64 {
	var maxWeight uint64
	for _, p := range parents {
		if p.GetWeight() > maxWeight {
			maxWeight = p.GetWeight()
		}
	}
	return maxWeight + 1
}

func (t Tx) Compare(tx Txi) bool {
	switch tx := tx.(type) {
	case *Tx:
		if t.GetTxHash().Cmp(tx.GetTxHash()) == 0 {
			return true
		}
		return false
	default:
		return false
	}
}