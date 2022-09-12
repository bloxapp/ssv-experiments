package ecies

import (
	"crypto/rand"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func EncryptECIES(plain []byte, pubKey []byte) []byte {
	pkECDSA, _ := crypto.DecompressPubkey(pubKey)
	pk := ecies.ImportECDSAPublic(pkECDSA)
	ct, _ := ecies.Encrypt(rand.Reader, pk, plain, nil, nil)
	return ct
}

func privKeyECIESFromByts(privKey []byte) *ecies.PrivateKey {
	n := &big.Int{}
	n.SetBytes(privKey)
	skECDSA, _ := crypto.ToECDSA(privKey)
	return ecies.ImportECDSA(skECDSA)
}

func DecryptECIES(ct []byte, privKey []byte) []byte {
	sk := privKeyECIESFromByts(privKey)
	pt, _ := sk.Decrypt(ct, nil, nil)
	return pt
}

func TestEncryptDecrypt(t *testing.T) {
	_ = bls.Init(bls.BLS12_381)
	_ = bls.SetETHmode(bls.EthModeDraft07)

	prv2, err := ecies.GenerateKey(rand.Reader, crypto.S256(), nil)
	if err != nil {
		t.Fatal(err)
	}
	sk2Byts := prv2.ExportECDSA().D.Bytes()
	pk2Bytes := crypto.CompressPubkey(&prv2.ExportECDSA().PublicKey)

	share := bls.SecretKey{}
	share.SetByCSPRNG()
	shareByts := share.Serialize()

	ct := EncryptECIES(shareByts, pk2Bytes)
	plainDecrypted := DecryptECIES(ct, sk2Byts)

	require.EqualValues(t, shareByts, plainDecrypted)
}
