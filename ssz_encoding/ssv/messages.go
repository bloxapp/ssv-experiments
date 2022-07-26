package ssv

import (
	"ssv-experiments/ssz_encoding/qbft"
	"ssv-experiments/ssz_encoding/types"
)

//go:generate go run .../fastssz/sszgen --path . --include ../qbft,../types

type PartialSignature struct {
	Slot          uint64
	Signature     [96]byte `ssz-size:"96"`
	SigningRoot   [32]byte `ssz-size:"32"`
	Signer        uint64
	Justification *qbft.SignedMessageHeader
}

type SignedPartialSignatures struct {
	ID                types.MessageID     `ssz-size:"52"`
	PartialSignatures []*PartialSignature `ssz-max:"13"`
}
