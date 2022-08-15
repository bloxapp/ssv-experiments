package ssv_state

import (
	"ssv-experiments/ssz_encoding/qbft"
	"ssv-experiments/ssz_encoding/types"
)

type QBFTState struct {
	Share                           Share
	ID                              types.MessageID
	Round                           uint64
	Height                          uint64
	LastPreparedRound               uint64
	LastPreparedValue               *types.ConsensusInput
	ProposalAcceptedForCurrentRound *qbft.SignedMessage
	Decided                         bool
	DecidedValue                    *types.ConsensusInput
}
