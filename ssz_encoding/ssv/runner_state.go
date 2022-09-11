package ssv

import (
	"ssv-experiments/ssz_encoding/qbft"
	"ssv-experiments/ssz_encoding/types"
)

type PartialSigContainer []*PartialSignature

// Add adds partial signature to the container for unique signers ONLY.
func (c PartialSigContainer) Add(sigMsg *PartialSignature) error {
	panic("implement")
}

func (c PartialSigContainer) ReconstructSignature(root, validatorPubKey []byte) ([]byte, error) {
	panic("implement")
}

type State struct {
	PreConsensusContainer  PartialSigContainer `ssz-max:"13"`
	PostConsensusContainer PartialSigContainer `ssz-max:"13"`
	RunningInstance        *qbft.Instance
	DecidedValue           *types.ConsensusInput
	StartingDuty           types.Duty
	Finished               bool

	Share          types.Share
	QBFTController qbft.Controller
	BeaconNetwork  [4]byte `ssz-size:"4"`
	BeaconRole     [4]byte `ssz-size:"4"`
}
