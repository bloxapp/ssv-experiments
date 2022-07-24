package ssv

import (
	"ssv-experiments/ssz_encoding/qbft"
	"ssv-experiments/ssz_encoding/types"
)

//go:generate go run .../fastssz/sszgen --path . --include ../qbft,../types

type PartialSignature struct {
}

type SignedPartialSignatures struct {
	ID                types.MessageID     `ssz-size:"52"`
	PartialSignatures []*PartialSignature `ssz-max:"13"`
	Justification     *qbft.SignedMessageHeader
}

type SignedPartialSignatureHeader struct {
	ID                types.MessageID     `ssz-size:"52"`
	PartialSignatures []*PartialSignature `ssz-max:"13"`
	JustificationRoot [32]byte            `ssz-size:"32"`
}
