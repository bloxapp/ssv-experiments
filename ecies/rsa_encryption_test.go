package ecies

import (
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRSAEncryptDecrypt(t *testing.T) {
	_ = bls.Init(bls.BLS12_381)
	_ = bls.SetETHmode(bls.EthModeDraft07)

	_, pkByts, err := GenerateKey()
	require.NoError(t, err)
	pk, err := PemToPublicKey(pkByts)
	require.NoError(t, err)

	fmt.Printf("pk L %d bytes\n", len(pkByts))

	share := bls.SecretKey{}
	share.SetByCSPRNG()
	shareByts := share.Serialize()

	ct, err := Encrypt(pk, shareByts)
	require.NoError(t, err)

	fmt.Printf("cipher text L %d bytes\n", len(ct))

	fmt.Printf("cipher text compressed L %d bytes\n", len(ct))

}
