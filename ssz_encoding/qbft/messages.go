package qbft

import (
	"ssv-experiments/ssz_encoding/types"
)

//go:generate go run .../fastssz/sszgen --path . --include ../types

type BaseMessage struct {
	ID     types.MessageID `ssz-size:"52"`
	Height uint64
	Round  uint64
	Digest [32]byte `ssz-size:"32"`
}

type SignedCommitMessage struct {
	Message   *BaseMessage
	Signers   []uint64 `ssz-max:"13"`
	Signature [96]byte `ssz-size:"96"`
}

type SignedPrepareMessage struct {
	Message   *BaseMessage
	Signers   []uint64 `ssz-max:"13"`
	Signature [96]byte `ssz-size:"96"`
}

type SignedProposalMessage struct {
	Message   *BaseMessage
	Signers   []uint64 `ssz-max:"13"`
	Signature [96]byte `ssz-size:"96"`
}

type RoundChangeMessage struct {
	Message       *BaseMessage
	PreparedRound uint64
}

type SignedRoundChangeMessage struct {
	Message   *RoundChangeMessage
	Signers   []uint64 `ssz-max:"13"`
	Signature [96]byte `ssz-size:"96"`
}
