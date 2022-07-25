package qbft

import (
	"github.com/stretchr/testify/require"
	"ssv-experiments/ssz_encoding/types"
	"testing"
)

var id = types.MessageID{0x1, 0x2, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x2, 0x3, 0x0, 0x0, 0x0, 0x0}
var prepareMsg = &SignedMessageHeader{
	Message: MessageHeader{
		ID:            id,
		Type:          Prepare,
		Height:        2,
		Round:         2,
		InputRoot:     [32]byte{},
		PreparedRound: 1,
	},
}
var roundChangeMsg = &SignedMessage{
	Message: Message{
		ID:     id,
		Type:   RoundChange,
		Height: 2,
		Round:  2,
		Input: types.ConsensusInput{ // input is used as prepared value
			Duty: types.Duty{
				Type:   types.BNRoleAttester,
				PubKey: [48]byte{},
			},
			DataSSZ: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
		PreparedRound: 1,
	},
	RoundChangeJustifications: []*SignedMessageHeader{
		prepareMsg,
		prepareMsg,
		prepareMsg, // etc.
	},
}
var roundChangeMsgHeader = func() *SignedMessageHeader {
	ret, _ := roundChangeMsg.ToSignedMessageHeader()
	return ret
}()

var proposalMsg = &SignedMessage{
	Message: Message{
		ID:     id,
		Type:   Proposal,
		Height: 2,
		Round:  2,
		Input: types.ConsensusInput{ // input is used as prepared value
			Duty: types.Duty{
				Type:   types.BNRoleAttester,
				PubKey: [48]byte{},
			},
			DataSSZ: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		},
		PreparedRound: 1,
	},
	RoundChangeJustifications: []*SignedMessageHeader{
		prepareMsg,
		prepareMsg,
		prepareMsg, // etc.
	},
	ProposalJustifications: []*SignedMessageHeader{
		roundChangeMsgHeader,
		roundChangeMsgHeader,
		roundChangeMsgHeader, // etc.
	},
}

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

func TestRoundChangeMsg(t *testing.T) {

}
