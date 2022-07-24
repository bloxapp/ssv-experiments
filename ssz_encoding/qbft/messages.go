package qbft

//go:generate go run .../fastssz/sszgen --path .

type CommitMessage struct {
	Height uint64
	Round  uint64
	Digest [32]byte `ssz-size:"32"`
}

type SignedCommitMessage struct {
	Message   *CommitMessage
	Signers   []uint64 `ssz-max:"13"`
	Signature [96]byte `ssz-size:"96"`
}

type PrepareMessage struct {
	Height uint64
	Round  uint64
	Digest [32]byte `ssz-size:"32"`
}

type SignedPrepareMessage struct {
	Message   *PrepareMessage
	Signers   []uint64 `ssz-max:"13"`
	Signature [96]byte `ssz-size:"96"`
}

type ProposalMessage struct {
	Height uint64
	Round  uint64
	Digest [32]byte `ssz-size:"32"`
}

type SignedProposalMessage struct {
	Message   *ProposalMessage
	Signers   []uint64 `ssz-max:"13"`
	Signature [96]byte `ssz-size:"96"`
}

type RoundChangeMessage struct {
	Height        uint64
	Round         uint64
	PreparedValue [32]byte `ssz-size:"32"`
	PreparedRound uint64
}

type SignedRoundChangeMessage struct {
	Message   *RoundChangeMessage
	Signers   []uint64 `ssz-max:"13"`
	Signature [96]byte `ssz-size:"96"`
}
