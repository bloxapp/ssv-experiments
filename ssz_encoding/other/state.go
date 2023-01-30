package other

import (
	"ssv-experiments/ssz_encoding/ssv"
	"ssv-experiments/ssz_encoding/types"
)

type PartialSigContainer []*ssv.PartialSignature

// Add adds partial signature to the container for unique signers ONLY.
func (c PartialSigContainer) Add(sigMsg *ssv.PartialSignature) error {
	panic("implement")
}

func (c PartialSigContainer) ReconstructSignature(root, validatorPubKey []byte) ([]byte, error) {
	panic("implement")
}

type State struct {
	PreConsensusContainer  PartialSigContainer `ssz-max:"13"`
	PostConsensusContainer PartialSigContainer `ssz-max:"13"`
	RunningInstance        *Instance
	DecidedValue           *types.ConsensusData
	StartingDuty           types.Duty
	Finished               bool
}
