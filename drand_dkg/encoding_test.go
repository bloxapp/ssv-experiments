package drand_dkg

import (
	"encoding/hex"
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/stretchr/testify/require"
	"testing"
)

func skFromString(t *testing.T, s string) bls.SecretKey {
	ret := bls.SecretKey{}
	byts, _ := hex.DecodeString(s)
	require.NoError(t, ret.Deserialize(byts))
	return ret
}

func blsID(t *testing.T, id int) bls.ID {
	blsID := bls.ID{}
	require.NoError(t, blsID.SetDecString(fmt.Sprintf("%d", id)))
	return blsID
}

// tested against the below encoding from drand
/*
// taken from https://github.com/alonmuroch/kyber/blob/bls-dkg/share/dkg/dkg_test.go#L126-L173
	bytssk, _ := share.V.MarshalBinary()
	sk := &bls2.SecretKey{}
	require.NoError(t, sk.Deserialize(bytssk))
	fmt.Printf("sk: %s\n", hex.EncodeToString(sk.Serialize()))

	byts, _ := pubShare.V.MarshalBinary()
	fmt.Printf("pubshare: %x\n", byts)
	pk := &bls2.PublicKey{}
	require.NoError(t, pk.Deserialize(byts))
	fmt.Printf("pk hex: %s\n", pk.GetHexString())
*/
func TestEncodeFromDrand(t *testing.T) {
	_ = bls.Init(bls.BLS12_381)
	_ = bls.SetETHmode(bls.EthModeDraft07)

	sk1 := skFromString(t, "7332e60a364cc722dc6b4f992f94e97ca6a458c76afe6add26f38040e4fb32e6")
	sk2 := skFromString(t, "58004daa4654b306f4a39f8cdc42c8b7d1ad25b40837820f451e0802b98149db")
	sk3 := skFromString(t, "6d0dc2ef489735ba5c9515880458121d83881688908bd577562a8ce8406ce823")
	sk4 := skFromString(t, "3e6d9e861376d1f4e105d9829e32eda86877874203fd09165a190ef279be0dbd")
	validatorSK := skFromString(t, "4ab7e4bbeee1f4c5e0b24da4f4ac9c66aeb00bbfb8e233e1fbaaf5a3c2daa343")
	//validatorPK := "8eac72c83c7e416fab9cc1933c5b73702d4fbf83819738d12af33331ed85afb3df99f43502697861919e7c40daa2e93d"

	s := bls.SecretKey{}
	require.NoError(t, s.Recover(
		[]bls.SecretKey{sk1, sk2, sk3, sk4},
		[]bls.ID{blsID(t, 1), blsID(t, 2), blsID(t, 3), blsID(t, 4)},
	))

	require.EqualValues(t, validatorSK.GetHexString(), s.GetHexString())
}
