package ssv_state

import (
	"bytes"
	"github.com/pkg/errors"
	ssvstruct "ssv-experiments/ssz_encoding/ssv"
	ssvtypes "ssv-experiments/ssz_encoding/types"
)

//go:generate go run .../fastssz/sszgen --path . --include ../ssz_encoding/ssv,../ssz_encoding/types

type State struct {
	PrePartialSignatures  []*ssvstruct.PartialSignature `ssz-max:"128"` // TODO why 128?
	PostPartialSignatures []*ssvstruct.PartialSignature `ssz-max:"128"`
	RunningConsensus      *QBFTInstance
}

func (s *State) AddPreConsensusPartialSig(partialSig *ssvstruct.PartialSignature) bool {
	return s.addUniquePartialSignature(s.PrePartialSignatures, partialSig)
}

func (s *State) AddPostConsensusPartialSig(partialSig *ssvstruct.PartialSignature) bool {
	return s.addUniquePartialSignature(s.PostPartialSignatures, partialSig)
}

func (s *State) Reconstruct(root [32]byte) ([96]byte, error) {
	msgs := make([]*ssvstruct.PartialSignature, 0)
	for _, msg := range s.PostPartialSignatures {
		if bytes.Equal(root[:], msg.SigningRoot[:]) {
			msgs = append(msgs, msg)
		}
	}
	return reconstructSignature(msgs)
}

func (s *State) DecidedValue() (*ssvtypes.ConsensusInput, error) {
	if !s.IsDecided() {
		return nil, nil
	}

	bytsSSZ := s.RunningConsensus.DecidedValue()
	ret := &ssvtypes.ConsensusInput{}
	if err := ret.UnmarshalSSZ(bytsSSZ); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal consensus input")
	}
	return ret, nil
}

func (s *State) IsDecided() bool {
	if s.RunningConsensus == nil {
		return false
	}
	return s.RunningConsensus.IsDecided()
}

// addUniquePartialSignature returns true if added (unique by signer and root)
func (s *State) addUniquePartialSignature(container []*ssvstruct.PartialSignature, partialSig *ssvstruct.PartialSignature) bool {
	for _, msg := range container {
		if msg.Signer == partialSig.Signer && bytes.Equal(msg.SigningRoot[:], partialSig.SigningRoot[:]) {
			return false
		}
	}

	container = append(container, partialSig)
	return true
}
