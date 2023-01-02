package drand_dkg

import (
	"github.com/drand/kyber/share/dkg"
	"github.com/pkg/errors"
)

type Node struct {
	Index uint32
	drand *dkg.DistKeyGenerator
	ecies *ECIES
}

func NewNode(index uint32, s dkg.Suite) *Node {
	return &Node{
		Index: index,
		ecies: NewRandomECIES(s),
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
