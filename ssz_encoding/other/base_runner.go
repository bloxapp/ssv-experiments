package other

import (
	"ssv-experiments/ssz_encoding/types"
)

type BaseRunner struct {
	State          *State
	Share          types.Share
	QBFTController Controller
	BeaconNetwork  [4]byte `ssz-size:"4"`
	BeaconRole     [4]byte `ssz-size:"4"`
}
