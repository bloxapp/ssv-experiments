package types

type Operator struct {
	OperatorID uint64
	PubKey     [48]byte `ssz-size:"48"`
}

type Share struct {
	OperatorID            uint64
	ValidatorPubKey       [48]byte    `ssz-size:"48"`
	SharePubKey           [48]byte    `ssz-size:"48"`
	Committee             []*Operator `ssz-max:"13"`
	Quorum, PartialQuorum uint64
	DomainType            DomainType `ssz-size:"4"`
	Graffiti              []byte     `ssz-size:"32"`
}
