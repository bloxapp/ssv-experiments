package ssv

//type PartialSigMsgType uint64
//
//const (
//	// PostConsensusPartialSig is a partial signature over a decided duty (attestation data, block, etc)
//	PostConsensusPartialSig PartialSigMsgType = iota
//	// RandaoPartialSig is a partial signature over randao reveal
//	RandaoPartialSig
//	// SelectionProofPartialSig is a partial signature for aggregator selection proof
//	SelectionProofPartialSig
//	// ContributionProofs is the partial selection proofs for sync committee contributions (it's an array of sigs)
//	ContributionProofs
//)

type PartialSignature struct {
	Signature   [96]byte `ssz-size:"96"`
	SigningRoot [32]byte `ssz-size:"32"`
	Signer      uint64
}

type PartialSignatures struct {
	Type     uint64
	Messages []*PartialSignature `ssz-max:"13"`
}

type SignedPartialSignatures struct {
	Message   PartialSignatures
	Signature [96]byte `ssz-size:"96"`
	Signer    uint64
}
