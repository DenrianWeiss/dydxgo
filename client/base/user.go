package base

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type BaseUser struct {
	BaseClient
	Address         common.Address
	StarkPublicKey  []byte
	StarkPrivateKey []byte
}

func (b *BaseUser) StarkKeyToUint256() *big.Int {
	// todo
	return big.NewInt(0)
}
