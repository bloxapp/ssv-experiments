package types

import (
	bellatrix2 "github.com/attestantio/go-eth2-client/api/v1/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/prysmaticlabs/go-bitfield"
)

type Contribution struct {
	K [96]byte `ssz-size:"96"`
	V *altair.SyncCommitteeContribution
}

type ConsensusData struct {
	Duty                   Duty
	AttestationData        *phase0.AttestationData
	BlockData              *bellatrix.BeaconBlock
	BlindedBlockData       *bellatrix2.BlindedBeaconBlock
	AggregateAndProof      *phase0.AggregateAndProof
	SyncCommitteeBlockRoot [32]byte `ssz-size:"32"`
	// SyncCommitteeContribution map holds as key the selection proof for the contribution
	SyncCommitteeContribution []*Contribution `ssz-max:"13"`
}

// NewConsensusData is a special constructor to avoid go-eth2-client bug with !len(32) block hash bug
func NewConsensusData(Duty Duty) *ConsensusData {
	beaconBlockEmpty := new(bellatrix.BeaconBlock)
	beaconBlockEmpty.Body = new(bellatrix.BeaconBlockBody)
	beaconBlockEmpty.Body.ETH1Data = new(phase0.ETH1Data)
	beaconBlockEmpty.Body.ETH1Data.BlockHash = make([]byte, 32)
	beaconBlockEmpty.Body.SyncAggregate = new(altair.SyncAggregate)
	beaconBlockEmpty.Body.SyncAggregate.SyncCommitteeBits = bitfield.NewBitvector512()

	beaconBlindedBlockEmpty := new(bellatrix2.BlindedBeaconBlock)
	beaconBlindedBlockEmpty.Body = new(bellatrix2.BlindedBeaconBlockBody)
	beaconBlindedBlockEmpty.Body.ETH1Data = new(phase0.ETH1Data)
	beaconBlindedBlockEmpty.Body.ETH1Data.BlockHash = make([]byte, 32)
	beaconBlindedBlockEmpty.Body.SyncAggregate = new(altair.SyncAggregate)
	beaconBlindedBlockEmpty.Body.SyncAggregate.SyncCommitteeBits = bitfield.NewBitvector512()

	return &ConsensusData{
		Duty:             Duty,
		BlockData:        beaconBlockEmpty,
		BlindedBlockData: beaconBlindedBlockEmpty,
	}
}

//func (ci *ConsensusData) GetAttestationData() (*phase0.AttestationData, error) {
//	ret := &phase0.AttestationData{}
//	if err := ret.UnmarshalSSZ(ci.DataSSZ); err != nil {
//		return nil, errors.Wrap(err, "could not unmarshal ssz")
//	}
//	return ret, nil
//}
//
//func (ci *ConsensusData) GetBlockData() (*bellatrix.BeaconBlock, error) {
//	ret := &bellatrix.BeaconBlock{}
//	if err := ret.UnmarshalSSZ(ci.DataSSZ); err != nil {
//		return nil, errors.Wrap(err, "could not unmarshal ssz")
//	}
//	return ret, nil
//}
//
//func (ci *ConsensusData) GetBlindedBlockData() (*v1.BlindedBeaconBlock, error) {
//	ret := &v1.BlindedBeaconBlock{}
//	if err := ret.UnmarshalSSZ(ci.DataSSZ); err != nil {
//		return nil, errors.Wrap(err, "could not unmarshal ssz")
//	}
//	return ret, nil
//}
//
//func (ci *ConsensusData) GetAggregateAndProof() (*phase0.AggregateAndProof, error) {
//	ret := &phase0.AggregateAndProof{}
//	if err := ret.UnmarshalSSZ(ci.DataSSZ); err != nil {
//		return nil, errors.Wrap(err, "could not unmarshal ssz")
//	}
//	return ret, nil
//}
//
//func (ci *ConsensusData) GetSyncCommitteeBlockRoot() (phase0.Root, error) {
//	if len(ci.DataSSZ) < 32 {
//		return phase0.Root{}, errors.New("could not unmarshal ssz")
//	}
//	ret := phase0.Root{}
//	copy(ret[:], ci.DataSSZ[0:32])
//	return ret, nil
//}
//
////func (ci *ConsensusData) GetContributionMap() (map[phase0.BLSSignature]*altair.SyncCommitteeContribution, error) {
////	mapping := &ContributionMapping{}
////	if err := mapping.UnmarshalSSZ(ci.DataSSZ); err != nil {
////		return nil, errors.Wrap(err, "could not unmarshal ssz")
////	}
////	return ret, nil
////}
