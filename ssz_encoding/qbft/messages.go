package qbft

import (
	"github.com/pkg/errors"
	"ssv-experiments/ssz_encoding/types"
)

//go:generate go run .../fastssz/sszgen --path . --include ../types

// Message includes the full consensus input to be decided on, used for proposal and round-change messages
type Message struct {
	ID     types.MessageID `ssz-size:"52"`
	Height uint64
	Round  uint64
	Input  types.ConsensusInput
	// PreparedRound an optional field used for round-change
	PreparedRound uint64
}

func (msg Message) ToMessageHeader() (MessageHeader, error) {
	r, err := msg.Input.HashTreeRoot()
	if err != nil {
		return MessageHeader{}, errors.Wrap(err, "failed to get input root")
	}
	return MessageHeader{
		ID:            msg.ID,
		Height:        msg.Height,
		Round:         msg.Round,
		InputRoot:     r,
		PreparedRound: msg.PreparedRound,
	}, nil
}

// SignedMessage includes a signature over Message AND optional justification fields (not signed over)
type SignedMessage struct {
	Message   Message
	Signers   []uint64 `ssz-max:"13"`
	Signature [96]byte `ssz-size:"96"`

	RoundChangeJustifications []*SignedMessageHeader `ssz-max:"13"`
	ProposalJustifications    []*SignedMessageHeader `ssz-max:"13"`
}

func (msg *SignedMessage) ToSignedMessageHeader() (*SignedMessageHeader, error) {
	header, err := msg.Message.ToMessageHeader()
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert to header")
	}

	return &SignedMessageHeader{
		Message:   header,
		Signers:   msg.Signers,
		Signature: msg.Signature,
	}, nil
}

// MessageHeader includes just the root of the input to be decided on (to save space), used for prepare and commit messages
type MessageHeader struct {
	ID            types.MessageID `ssz-size:"52"`
	Height        uint64
	Round         uint64
	InputRoot     [32]byte `ssz-size:"32"`
	PreparedRound uint64
}

// SignedMessageHeader includes a signature over MessageHeader
type SignedMessageHeader struct {
	Message   MessageHeader
	Signers   []uint64 `ssz-max:"13"`
	Signature [96]byte `ssz-size:"96"`
}
