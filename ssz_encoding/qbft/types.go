package qbft

import "ssv-experiments/ssz_encoding/types"

// MsgContainer holds all accepted messages
type MsgContainer []*SignedMessage

func (c MsgContainer) PerRound(round uint64) []SignedMessage {
	panic("implement")
}

func (c MsgContainer) PerRoundAndValue(round uint64, value []byte) []SignedMessage {
	panic("implement")
}

func (c MsgContainer) LongestUniqueSignersForRoundAndValue(round uint64, value []byte) []SignedMessage {
	panic("implement")
}

type State struct {
	Share                           types.Share
	ID                              [32]byte `ssz-size:"32"`
	Round                           uint64
	Height                          uint64
	LastPreparedRound               uint64
	LastPreparedValue               *types.ConsensusInput
	ProposalAcceptedForCurrentRound *SignedMessage
	Decided                         bool
	DecidedValue                    *types.ConsensusInput

	ProposeContainer     MsgContainer `ssz-max:"256"` // TODO - why 256 max per instance?
	PrepareContainer     MsgContainer `ssz-max:"256"`
	CommitContainer      MsgContainer `ssz-max:"256"`
	RoundChangeContainer MsgContainer `ssz-max:"256"`
}

type Instance struct {
	State      State
	StartValue types.ConsensusInput
}

// FutureMsgContainer holds for each operator (by order in the share) the highest height msg received and validated.
// Comment: Every decided instance the controller should check if the decided height is bigger than the stored height in the container, set the highest of the two
type FutureMsgContainer []uint64

type Controller struct {
	ID                 [32]byte `ssz-size:"32"`
	Height             uint64
	ActiveInstances    []*Instance        `ssz-max:"5"`
	FutureMsgContainer FutureMsgContainer `ssz-max:"13"`
	Domain             types.DomainType   `ssz-size:"4"`
	Share              types.Share
}
