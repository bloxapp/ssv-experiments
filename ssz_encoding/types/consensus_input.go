package types

import (
	v1 "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
)

//type ContributionMapping struct {
//	Keys   [][96]byte `ssz-max:"13"`
//	Values [][]byte   `ssz-max:"13,2048"`
//}

type ConsensusInput struct {
	Duty    Duty
	DataSSZ []byte `ssz-max:"2048"`
}

func (ci *ConsensusInput) GetAttestationData() (*phase0.AttestationData, error) {
	ret := &phase0.AttestationData{}
	if err := ret.UnmarshalSSZ(ci.DataSSZ); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal ssz")
	}
	return ret, nil
}

func (ci *ConsensusInput) GetBlockData() (*bellatrix.BeaconBlock, error) {
	ret := &bellatrix.BeaconBlock{}
	if err := ret.UnmarshalSSZ(ci.DataSSZ); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal ssz")
	}
	return ret, nil
}

func (ci *ConsensusInput) GetBlindedBlockData() (*v1.BlindedBeaconBlock, error) {
	ret := &v1.BlindedBeaconBlock{}
	if err := ret.UnmarshalSSZ(ci.DataSSZ); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal ssz")
	}
	return ret, nil
}

func (ci *ConsensusInput) GetAggregateAndProof() (*phase0.AggregateAndProof, error) {
	ret := &phase0.AggregateAndProof{}
	if err := ret.UnmarshalSSZ(ci.DataSSZ); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal ssz")
	}
	return ret, nil
}

func (ci *ConsensusInput) GetSyncCommitteeBlockRoot() (phase0.Root, error) {
	if len(ci.DataSSZ) < 32 {
		return phase0.Root{}, errors.New("could not unmarshal ssz")
	}
	ret := phase0.Root{}
	copy(ret[:], ci.DataSSZ[0:32])
	return ret, nil
}

//func (ci *ConsensusInput) GetContributionMap() (map[phase0.BLSSignature]*altair.SyncCommitteeContribution, error) {
//	mapping := &ContributionMapping{}
//	if err := mapping.UnmarshalSSZ(ci.DataSSZ); err != nil {
//		return nil, errors.Wrap(err, "could not unmarshal ssz")
//	}
//	return ret, nil
//}
