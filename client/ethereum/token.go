package ethereum

import (
	"github.com/denrianweiss/dydxgo/client/ethereum/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

var tokenAddress = map[int64]string{
	1:        "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
	3:        "0x8707a5bf4c2842d46b31a405ba41b858c0f876c4",
	5:        "0xF7a2fa2c2025fFe64427dd40Dc190d47ecC8B36e",
	11155111: "0x7fC9C132268E0E414991449c003DbdB3E73E2059",
}

type Token struct {
	ETHClient
	token *abi.Erc20
}

func (t *Token) New(e ETHClient) *Token {
	t.ETHClient = e
	contract, _ := abi.NewErc20(common.HexToAddress(tokenAddress[t.NetworkId]), t.ethClient)
	t.token = contract
	return t
}

func (t *Token) SetAllowance(amount *big.Int, transact *bind.TransactOpts) (*types.Transaction, error) {
	return t.token.Approve(transact, common.HexToAddress(exchangeAddress[t.NetworkId]), amount)
}

func (t *Token) GetAllowance() (*big.Int, error) {
	return t.token.Allowance(&bind.CallOpts{}, t.Address, common.HexToAddress(exchangeAddress[t.NetworkId]))
}

func (t *Token) GetContract() *abi.Erc20 {
	return t.token
}
