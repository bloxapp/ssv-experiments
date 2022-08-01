package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMsgID(t *testing.T) {
	t.Run("validator index", func(t *testing.T) {
		id := NewMsgIDValidator(12, BNRoleAttester, ConsensusProposeMsgType)
		require.EqualValues(t, 12, id.GetValidatorIndex())
		require.EqualValues(t, BNRoleAttester, id.GetRoleType())
		require.EqualValues(t, ConsensusProposeMsgType, id.GetMsgType())
	})

	t.Run("dkg index", func(t *testing.T) {
		address := common.Address{0x99, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x3, 0x1, 0x1, 0x1, 0x1, 0x1, 0x5, 0x1, 0x1, 0x1, 0x11}
		id := NewMsgIDETHAddress(address, 133, DKGInitMsgType)
		require.EqualValues(t, address, id.GetETHAddress())
		require.EqualValues(t, 133, id.GetDKGIndex())
		require.EqualValues(t, DKGInitMsgType, id.GetMsgType())
	})
}
