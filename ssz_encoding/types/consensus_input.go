package types

//go:generate go run .../fastssz/sszgen --path . --exclude-objs MessageID,MessageBytes

type ConsensusInput struct {
}
