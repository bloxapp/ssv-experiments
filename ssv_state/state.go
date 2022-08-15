package ssv_state

import (
	"bytes"
	ssvstruct "ssv-experiments/ssz_encoding/ssv"
)

//go:generate go run .../fastssz/sszgen --path . --include ../ssz_encoding/ssv,../ssz_encoding/types

type State struct {
	PrePartialSignatures  []*ssvstruct.PartialSignature `ssz-max:"128"` // TODO why 128?
	PostPartialSignatures []*ssvstruct.PartialSignature `ssz-max:"128"`
	// RunningConsensusState is a pointer so it can be changed by the instance
	RunningConsensusState *QBFTState
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
