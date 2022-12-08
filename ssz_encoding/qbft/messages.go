package qbft

import (
	"bytes"
	"github.com/pkg/errors"
	"ssv-experiments/ssz_encoding/types"
)

// Message includes the full consensus input to be decided on, used for decided, proposal and round-change messages
type Message struct {
	Height uint64
	Round  uint64
	Root   []byte `ssz-size:"32"`
	// PreparedRound an optional field used for round-change
	PreparedRound uint64
}

// SignedMessage includes a signature over Message AND optional justification fields (not signed over)
type SignedMessage struct {
	Message   Message
	Signers   []uint64 `ssz-max:"13"`
	Signature [96]byte `ssz-size:"96"`

	Justifications *Justifications
	Object         *types.ConsensusInput
}

func (msg *SignedMessage) Validate() error {
	if msg.Object != nil {
		r, err := msg.Object.HashTreeRoot()
		if err != nil {
			return errors.Wrap(err, "could not get object root")
		}
		if !bytes.Equal(msg.Message.Root[:], r[:]) {
			return errors.Wrap(err, "object root not equal to message root")
		}
	}
	return nil
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
	ret := make([]*SignedMessage, len(data))
	for i, byts := range data {
		msg := &SignedMessage{}
		if err := msg.UnmarshalSSZ(byts); err != nil {
			return nil, errors.Wrap(err, "could not unmarshal signed message")
		}
		ret[i] = msg
	}
	return ret, nil
}
