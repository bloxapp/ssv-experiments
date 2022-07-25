package types

//go:generate go run .../fastssz/sszgen --path . --exclude-objs MessageID,MessageBytes

type ConsensusInput struct {
	Duty    Duty
	DataSSZ []byte `ssz-max:"2048"`
}
