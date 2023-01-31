package types

import (
	bellatrix2 "github.com/attestantio/go-eth2-client/api/v1/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
)

type Contribution struct {
	K [96]byte `ssz-size:"96"`
	V *altair.SyncCommitteeContribution
}

type ConsensusData struct {
	Duty                   Duty
	DataSSZ                []byte   `ssz-max:"16384"` // 2^14
	SyncCommitteeBlockRoot [32]byte `ssz-size:"32"`
	// SyncCommitteeContribution map holds as key the selection proof for the contribution
	SyncCommitteeContribution []*Contribution `ssz-max:"13"`

	attestationData   *phase0.AttestationData
	blockData         *bellatrix.BeaconBlock
	blindedBlockData  *bellatrix2.BlindedBeaconBlock
	aggregateAndProof *phase0.AggregateAndProof
}

//// NewConsensusData is a special constructor to avoid go-eth2-client bug with !len(32) block hash bug
//func NewConsensusData(Duty Duty) *ConsensusData {
//	beaconBlockEmpty := new(bellatrix.BeaconBlock)
//	beaconBlockEmpty.Body = new(bellatrix.BeaconBlockBody)
//	beaconBlockEmpty.Body.ETH1Data = new(phase0.ETH1Data)
//	beaconBlockEmpty.Body.ETH1Data.BlockHash = make([]byte, 32)
//	beaconBlockEmpty.Body.SyncAggregate = new(altair.SyncAggregate)
//	beaconBlockEmpty.Body.SyncAggregate.SyncCommitteeBits = bitfield.NewBitvector512()
//
//	beaconBlindedBlockEmpty := new(bellatrix2.BlindedBeaconBlock)
//	beaconBlindedBlockEmpty.Body = new(bellatrix2.BlindedBeaconBlockBody)
//	beaconBlindedBlockEmpty.Body.ETH1Data = new(phase0.ETH1Data)
//	beaconBlindedBlockEmpty.Body.ETH1Data.BlockHash = make([]byte, 32)
//	beaconBlindedBlockEmpty.Body.SyncAggregate = new(altair.SyncAggregate)
//	beaconBlindedBlockEmpty.Body.SyncAggregate.SyncCommitteeBits = bitfield.NewBitvector512()
//
//	return &ConsensusData{
//		Duty:             Duty,
//		BlockData:        beaconBlockEmpty,
//		BlindedBlockData: beaconBlindedBlockEmpty,
//	}
//}

func (ci *ConsensusData) GetAttestationData() (*phase0.AttestationData, error) {
	if ci.attestationData != nil {
		return ci.attestationData, nil
	}
	ret := &phase0.AttestationData{}
	if err := ret.UnmarshalSSZ(ci.DataSSZ); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal ssz")
	}
	ci.attestationData = ret
	return ret, nil
}

func (ci *ConsensusData) GetBlockData() (*bellatrix.BeaconBlock, error) {
	if ci.blockData != nil {
		return ci.blockData, nil
	}

	ret := &bellatrix.BeaconBlock{}
	if err := ret.UnmarshalSSZ(ci.DataSSZ); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal ssz")
	}
	ci.blockData = ret
	return ret, nil
}

func (ci *ConsensusData) GetBlindedBlockData() (*bellatrix2.BlindedBeaconBlock, error) {
	if ci.blindedBlockData != nil {
		return ci.blindedBlockData, nil
	}

	ret := &bellatrix2.BlindedBeaconBlock{}
	if err := ret.UnmarshalSSZ(ci.DataSSZ); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal ssz")
	}
	ci.blindedBlockData = ret
	return ret, nil
}

func (ci *ConsensusData) GetAggregateAndProof() (*phase0.AggregateAndProof, error) {
	if ci.aggregateAndProof != nil {
		return ci.aggregateAndProof, nil
	}

	ret := &phase0.AggregateAndProof{}
	if err := ret.UnmarshalSSZ(ci.DataSSZ); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal ssz")
	}
	ci.aggregateAndProof = ret
	return ret, nil
}

func (ci *ConsensusData) GetSyncCommitteeBlockRoot() phase0.Root {
	return ci.SyncCommitteeBlockRoot
}

func (ci *ConsensusData) GetContributionMap() []*Contribution {
	return ci.SyncCommitteeContribution
}
