package ssv

import (
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/require"
	"ssv-experiments/ssz_encoding/qbft"
	"testing"
)

func TestPartialSignature_HashTreeRoot(t *testing.T) {
	signed := &SignedPartialSignatures{
		PartialSignatures: []*PartialSignature{
			{},
			{},
		},
		Justification: &qbft.SignedCommitMessage{
			Message: &qbft.CommitMessage{
				Height: 1,
				Round:  1,
				Digest: [32]byte{},
			},
			Signature: [96]byte{},
			Signers:   []uint64{1},
		},
	}

	rr, _ := signed.Justification.HashTreeRoot()
	fmt.Printf("%s\n", hex.EncodeToString(rr[:]))

	r, err := signed.HashTreeRoot()
	require.NoError(t, err)
	require.EqualValues(t, "734dabde959e9c520d86c6f67db913a1666c2d20022045f775779949b2447302", hex.EncodeToString(r[:]))
}

func TestSignedPartialSignatureHeader_HashTreeRoot(t *testing.T) {
	r, _ := hex.DecodeString("f6268696c94ea38d6526f6b05cee1ca04dd8708afd921ba2ad0d54c89aa4a26e")
	root := [32]byte{}
	copy(root[:], r[:])

	signed := &SignedPartialSignatureHeader{
		PartialSignatures: []*PartialSignature{
			{},
			{},
		},
		Root: root,
	}

	r2, err := signed.HashTreeRoot()
	require.NoError(t, err)
	require.EqualValues(t, "734dabde959e9c520d86c6f67db913a1666c2d20022045f775779949b2447302", hex.EncodeToString(r2[:]))
}
