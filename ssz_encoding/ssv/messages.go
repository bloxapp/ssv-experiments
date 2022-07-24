package ssv

import "ssv-experiments/ssz_encoding/qbft"

//go:generate go run .../fastssz/sszgen --path . --include ../qbft

type PartialSignature struct {
}

type SignedPartialSignatures struct {
	PartialSignatures []*PartialSignature `ssz-max:"13"`
	Justification     *qbft.SignedCommitMessage
}

type SignedPartialSignatureHeader struct {
	PartialSignatures []*PartialSignature `ssz-max:"13"`
	Root              [32]byte            `ssz-size:"32"`
}
