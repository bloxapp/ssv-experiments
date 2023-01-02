package drand_dkg

import (
	"fmt"
	bls "github.com/drand/kyber-bls12381"
	"github.com/drand/kyber/share/dkg"
	bls2 "github.com/drand/kyber/sign/bls"
	bls3 "github.com/herumi/bls-eth-go-binary/bls"
	"github.com/stretchr/testify/require"
	"ssv-experiments/ssz_encoding/types"
	"testing"
)

var TestSuite = bls.NewBLS12381Suite()
var TestAuthScheme = bls2.NewSchemeOnG1(TestSuite)

const (
	N = 4
	T = 3
)

var TestNodes = func() []*Node {
	return []*Node{
		NewNode(1, TestSuite.G1().(dkg.Suite)),
		NewNode(2, TestSuite.G1().(dkg.Suite)),
		NewNode(3, TestSuite.G1().(dkg.Suite)),
		NewNode(4, TestSuite.G1().(dkg.Suite)),
	}
}()

var TestDrandNodes = func() []dkg.Node {
	return []dkg.Node{
		{
			Index:  1,
			Public: TestNodes[0].ecies.Pub,
		},
		{
			Index:  2,
			Public: TestNodes[1].ecies.Pub,
		},
		{
			Index:  3,
			Public: TestNodes[2].ecies.Pub,
		},
		{
			Index:  4,
			Public: TestNodes[3].ecies.Pub,
		},
	}
}()

func reconstructSK(t *testing.T, sks []bls3.SecretKey) *bls3.SecretKey {
	s := &bls3.SecretKey{}
	require.NoError(t, s.Recover(
		sks,
		[]bls3.ID{blsID(t, 1), blsID(t, 2), blsID(t, 3), blsID(t, 4)},
	))
	return s
}

func TestDKGFull(t *testing.T) {
	types.InitBLS()

	nonce := GetNonce()
	for _, n := range TestNodes {
		require.NoError(t, n.SetupDrandWithConfig(&dkg.Config{
			Suite:     TestSuite.G1().(dkg.Suite),
			NewNodes:  TestDrandNodes,
			Threshold: T,
			Auth:      TestAuthScheme,
			Nonce:     nonce,
		}))
	}

	// Step 1
	var deals []*dkg.DealBundle
	for _, n := range TestNodes {
		d, err := n.drand.Deals()
		require.NoError(t, err)
		deals = append(deals, d)
	}
	require.NotEmpty(t, deals)

	// Step 2
	var respBundles []*dkg.ResponseBundle
	for _, n := range TestNodes {
		r, err := n.drand.ProcessDeals(deals)
		require.NoError(t, err)
		if r != nil {
			respBundles = append(respBundles, r)
		}
	}

	// Step 3
	var justifs []*dkg.JustificationBundle
	var results []*dkg.Result
	for _, n := range TestNodes {
		res, just, err := n.drand.ProcessResponses(respBundles)
		require.NoError(t, err)

		if res != nil {
			results = append(results, res)
		} else if just != nil {
			justifs = append(justifs, just)
		}
	}

	if len(justifs) == 0 {
		require.NotEmpty(t, results)

		var sks []bls3.SecretKey
		for i, res := range results {
			sk, err := resultToShareSecretKey(res)
			require.NoError(t, err)
			fmt.Printf("Index (%d): %x\n", i+1, sk.Serialize())
			sks = append(sks, *sk)
		}

		valSK := reconstructSK(t, sks)
		valPK, err := resultsToValidatorPK(results, TestSuite.G1().(dkg.Suite))
		require.NoError(t, err)
		require.EqualValues(t, valPK.Serialize(), valSK.GetPublicKey().Serialize())

		fmt.Printf("Validator SK: %x\nValidator PK: %x\n", valSK.Serialize(), valPK.Serialize())
	} else {
		for _, n := range TestNodes {
			res, err := n.drand.ProcessJustifications(justifs)
			require.NoError(t, err)
			require.NotNil(t, res)
			results = append(results, res)
		}
	}
}
