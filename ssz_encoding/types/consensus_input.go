package types

type ConsensusInput struct {
	Duty    Duty
	DataSSZ []byte `ssz-max:"2048"`
}
