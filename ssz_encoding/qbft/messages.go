package qbft

import (
	"ssv-experiments/ssz_encoding/types"
)

type Data struct {
	Root   []byte `ssz-size:"32"`
	Object *types.ConsensusInput
}

// Message includes the full consensus input to be decided on, used for decided, proposal and round-change messages
type Message struct {
	Height uint64
	Round  uint64
	Input  Data
	// PreparedRound an optional field used for round-change
	PreparedRound uint64
}

// SignedMessage includes a signature over Message AND optional justification fields (not signed over)
type SignedMessage struct {
	Message   Message
	Signers   []uint64 `ssz-max:"13"`
	Signature [96]byte `ssz-size:"96"`

	Justifications *Justifications
}

type Justifications struct {
	RoundChangeJustifications [][]byte `ssz-max:"13,1024"`
	ProposalJustifications    [][]byte `ssz-max:"13,1024"`
}

func (j *Justifications) GetRoundChangeJustifications() ([]*SignedMessage, error) {
	return j.toSignedMessages(j.RoundChangeJustifications)
}

func (j *Justifications) GetProposalJustifications() ([]*SignedMessage, error) {
	return j.toSignedMessages(j.ProposalJustifications)
}

func (j *Justifications) toSignedMessages(data [][]byte) ([]*SignedMessage, error) {
	panic("implement")
}
