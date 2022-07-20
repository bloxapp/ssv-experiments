package ssv

//go:generate go run .../fastssz/sszgen --path . [--objs PartialSignature,SignedPartialSignatures]

type PartialSignature struct {
}

type SignedPartialSignatures struct {
	Signatures *PartialSignature
}
