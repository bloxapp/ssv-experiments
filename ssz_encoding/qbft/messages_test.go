package qbft

import (
	"encoding/json"
	"fmt"
	"github.com/golang/snappy"
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

func TestSizeComparison(t *testing.T) {
	sszByts, _ := proposalMsg.MarshalSSZ()
	sszSnappyByts := snappy.Encode([]byte{}, sszByts)
	jsonByts, _ := json.Marshal(proposalMsg)
	jsonSnappyByts := snappy.Encode([]byte{}, jsonByts)

	fmt.Printf("ssz: %d\n", len(sszByts))
	fmt.Printf("ssz snappy: %d\n", len(sszSnappyByts))
	fmt.Printf("json: %d\n", len(jsonByts))
	fmt.Printf("json snappy: %d\n", len(jsonSnappyByts))
}
