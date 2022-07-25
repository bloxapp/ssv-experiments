package qbft

import (
	"github.com/stretchr/testify/require"
	"ssv-experiments/ssz_encoding/types"
	"testing"
)

func TestHashRoot(t *testing.T) {
	id := types.MessageID{0x1, 0x2, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x2, 0x3, 0x0, 0x0, 0x0, 0x0}
	signed := &SignedMessage{
		Message: Message{
			ID:     id,
			Height: 1,
			Round:  2,
			Input: types.ConsensusInput{
				Duty: types.Duty{
					Type:   types.BNRoleAttester,
					PubKey: [48]byte{},
				},
			},
			PreparedRound: 33,
		},
		Signature: [96]byte{},
		Signers:   []uint64{1},
		RoundChangeJustifications: []*SignedMessageHeader{
			{
				Message: MessageHeader{
					ID:            id,
					Height:        1,
					Round:         2,
					InputRoot:     [32]byte{},
					PreparedRound: 33,
				},
				Signature: [96]byte{},
				Signers:   []uint64{1},
			},
		},
		ProposalJustifications: []*SignedMessageHeader{
			{
				Message: MessageHeader{
					ID:            id,
					Height:        1,
					Round:         2,
					InputRoot:     [32]byte{},
					PreparedRound: 33,
				},
				Signature: [96]byte{},
				Signers:   []uint64{1},
			},
		},
	}

	r, err := signed.Message.HashTreeRoot()
	require.NoError(t, err)

	signedHeader, err := signed.ToSignedMessageHeader()
	require.NoError(t, err)
	r2, err := signedHeader.Message.HashTreeRoot()
	require.NoError(t, err)
	require.EqualValues(t, r, r2)
}
