package types

import "github.com/herumi/bls-eth-go-binary/bls"

type DomainType [4]byte

var (
	Shifu      = DomainType{0x0, 0x0, 0x0, 0x1}
	SSVMainnet = DomainType{0x1, 0x0, 0x0, 0x0}
)

func InitBLS() {
	_ = bls.Init(bls.BLS12_381)
	_ = bls.SetETHmode(bls.EthModeDraft07)
}
