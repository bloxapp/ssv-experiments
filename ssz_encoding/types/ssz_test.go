package types

import (
	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/prysmaticlabs/go-bitfield"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContributionsSSZ(t *testing.T) {
	t.Run("simple ssz marshaling", func(t *testing.T) {
		c := make(Contributions, 5)
		for i := 0; i < 5; i++ {
			c[i] = &Contribution{
				K: [96]byte{},
				V: &altair.SyncCommitteeContribution{
					Slot:              12,
					BeaconBlockRoot:   phase0.Root{},
					SubcommitteeIndex: 0,
					AggregationBits:   bitfield.NewBitvector128(),
					Signature:         phase0.BLSSignature{},
				},
			}
		}

		r1, err := c.HashTreeRoot()
		require.NoError(t, err)

		byts, err := c.MarshalSSZ()
		require.NoError(t, err)

		decodedC := Contributions{}
		require.NoError(t, decodedC.UnmarshalSSZ(byts))

		r2, err := decodedC.HashTreeRoot()
		require.NoError(t, err)

		require.EqualValues(t, r1, r2)
	})

	t.Run(">13 ssz marshaling", func(t *testing.T) {
		c := make(Contributions, 14)
		for i := 0; i < 14; i++ {
			c[i] = &Contribution{
				K: [96]byte{},
				V: &altair.SyncCommitteeContribution{
					Slot:              12,
					BeaconBlockRoot:   phase0.Root{},
					SubcommitteeIndex: 0,
					AggregationBits:   bitfield.NewBitvector128(),
					Signature:         phase0.BLSSignature{},
				},
			}
		}

		_, err := c.HashTreeRoot()
		require.EqualError(t, err, "incorrect list size")

		_, err = c.MarshalSSZ()
		require.EqualError(t, err, "ConsensusData.SyncCommitteeContribution (list length is higher than max value): max expected 13 and 14 found")
	})
}
