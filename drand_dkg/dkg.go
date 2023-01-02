package drand_dkg

import (
	"crypto/rsa"
	"github.com/drand/kyber/share/dkg"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/pkg/errors"
	"ssv-experiments/ssz_encoding/types"
)

type Node struct {
	Index         uint32
	drand         *dkg.DistKeyGenerator
	ecies         *ECIES
	EncryptionSK  *rsa.PrivateKey
	GenerateShare *bls.SecretKey
}

func NewNode(index uint32, s dkg.Suite) *Node {
	skByts, _, err := types.GenerateKey()
	if err != nil {
		panic("could not generate rsa key")
	}
	sk, _ := types.PemToPrivateKey(skByts)

	return &Node{
		Index:        index,
		ecies:        NewRandomECIES(s),
		EncryptionSK: sk,
	}
}

func (n *Node) SetupDrandWithConfig(c *dkg.Config) error {
	c.Longterm = n.ecies.Priv
	drand, err := dkg.NewDistKeyHandler(c)
	if err != nil {
		errors.Wrap(err, "could not generate drand DKG")
	}
	n.drand = drand
	return nil
}

func (n *Node) EncryptShare() ([]byte, error) {
	return types.Encrypt(&n.EncryptionSK.PublicKey, n.GenerateShare.Serialize())
}
