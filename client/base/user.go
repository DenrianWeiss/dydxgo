package base

import (
	"github.com/denrianweiss/dydxgo/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

type BaseUser struct {
	BaseClient
	Address           common.Address
	StarkPrivateKey   string
	ApiKeyCredentials *types.ApiKeyCredentials
}

func (b *BaseUser) StarkKeyToUint256() *big.Int {
	bI, _ := big.NewInt(0).SetString(strings.TrimPrefix(b.StarkPrivateKey, "0x"), 16)
	return bI
}
