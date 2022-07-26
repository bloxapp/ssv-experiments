package ssv

import (
	"ssv-experiments/ssz_encoding/qbft"
)

type PartialSigMsgType uint64

const (
	// PostConsensusPartialSig is a partial signature over a decided duty (attestation data, block, etc)
	PostConsensusPartialSig PartialSigMsgType = iota
	// RandaoPartialSig is a partial signature over randao reveal
	RandaoPartialSig
	// SelectionProofPartialSig is a partial signature for aggregator selection proof
	SelectionProofPartialSig
	// ContributionProofs is the partial selection proofs for sync committee contributions (it's an array of sigs)
	ContributionProofs
)

type PartialSignature struct {
	Slot        uint64
	Signature   [96]byte `ssz-size:"96"`
	SigningRoot [32]byte `ssz-size:"32"`
	// Justification is an optional param, the decided message post consensus
	Justification *qbft.SignedMessage
}

type PartialSignatures struct {
	Type              PartialSigMsgType
	PartialSignatures []*PartialSignature `ssz-max:"13"`
}

type SignedPartialSignatures struct {
	PartialSignatures PartialSignatures
	Signature         [96]byte `ssz-size:"96"`
	Signer            uint64
}
