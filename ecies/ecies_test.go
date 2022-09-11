package ecies

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/herumi/bls-eth-go-binary/bls"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	_ = bls.Init(bls.BLS12_381)
	_ = bls.SetETHmode(bls.EthModeDraft07)

	prv1, err := ecies.GenerateKey(rand.Reader, crypto.S256(), nil)
	if err != nil {
		t.Fatal(err)
	}

	prv2, err := ecies.GenerateKey(rand.Reader, crypto.S256(), nil)
	if err != nil {
		t.Fatal(err)
	}

	share := bls.SecretKey{}
	share.SetByCSPRNG()
	shareByts := share.Serialize()

	ct, err := ecies.Encrypt(rand.Reader, &prv2.PublicKey, shareByts, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("cipher text L %d bytes\n", len(ct))

	pt, err := prv2.Decrypt(ct, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(pt, shareByts) {
		t.Fatal("ecies: plaintext doesn't match message")
	}

	_, err = prv1.Decrypt(ct, nil, nil)
	if err == nil {
		t.Fatal("ecies: encryption should not have succeeded")
	}
}
