package drand_dkg

import (
	"github.com/drand/kyber"
	"github.com/drand/kyber/share"
	"github.com/drand/kyber/share/dkg"
	"github.com/drand/kyber/util/random"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/pkg/errors"
	"math/rand"
)

type ECIES struct {
	Priv kyber.Scalar
	Pub  kyber.Point
}

func NewRandomECIES(s dkg.Suite) *ECIES {
	private := s.Scalar().Pick(random.New())
	return &ECIES{
		Priv: private,
		Pub:  s.Point().Mul(private, nil),
	}
}

// NonceLength is the length of the nonce
const NonceLength = 32

// GetNonce returns a suitable nonce to feed in the DKG config.
func GetNonce() []byte {
	var nonce [NonceLength]byte
	n, err := rand.Read(nonce[:])
	if n != NonceLength {
		panic("could not read enough random bytes for nonce")
	}
	if err != nil {
		panic(err)
	}
	return nonce[:]
}

func resultToShareSecretKey(result *dkg.Result) (*bls.SecretKey, error) {
	share := result.Key.PriShare()
	bytsSk, err := share.V.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal share")
	}
	sk := &bls.SecretKey{}
	if err := sk.Deserialize(bytsSk); err != nil {
		return nil, errors.Wrap(err, "could not deserialized secret key")
	}
	return sk, nil
}

func resultsToValidatorPK(results []*dkg.Result, suite dkg.Suite) (*bls.PublicKey, error) {
	exp := share.NewPubPoly(suite, suite.Point().Base(), results[0].Key.Commitments())
	bytsPK, err := exp.Eval(0).V.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal share")
	}
	pk := &bls.PublicKey{}
	if err := pk.Deserialize(bytsPK); err != nil {
		return nil, errors.Wrap(err, "could not deserialized public key")
	}
	return pk, nil
}
