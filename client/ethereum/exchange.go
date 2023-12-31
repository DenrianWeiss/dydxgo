package ethereum

import (
	"github.com/denrianweiss/dydxgo/client/ethereum/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

var exchangeAddress = map[int64]string{
	1:        "0xD54f502e184B6B739d7D27a6410a67dc462D69c8",
	3:        "0x014F738EAd8Ec6C50BCD456a971F8B84Cd693BBe",
	5:        "0xFE76edf35648Cc733d57200646cb1Dc63d05462F",
	11155111: "0x3D05aaCd0fED84f65dE0D91e4621298E702911E2",
}

var tokenId = map[int64]string{
	1:        "0x02893294412a4c8f915f75892b395ebbf6859ec246ec365c3b1f56f47c3a0a5d",
	5:        "0x03bda2b4764039f2df44a00a9cf1d1569a83f95406a983ce4beb95791c376008",
	11155111: "0x01e70c509c4c6bfafe8b73d2fc1819444b2c0b435d4b82c0f24addff9565ce25",
}

type Exchange struct {
	ETHClient
	exchange *abi.Exchange
}

func (e *Exchange) New(c ETHClient) *Exchange {
	e.ETHClient = c
	exchangeAddr := exchangeAddress[c.NetworkId]
	contract, err := abi.NewExchange(common.HexToAddress(exchangeAddr), e.ethClient)
	if err != nil {
		panic(err)
	}
	e.exchange = contract
	return e
}

func (e *Exchange) Deposit(amount *big.Int, publicKey string, positionId *big.Int, transact *bind.TransactOpts) (*types.Transaction, error) {
	if transact.Signer == nil {
		transact.Signer = e.ethSigner
	}
	tokenIdUint, _ := big.NewInt(0).SetString(strings.TrimPrefix(tokenId[e.NetworkId], "0x"), 16)
	return e.exchange.Deposit0(transact, e.PublicKeyToUint256(publicKey), tokenIdUint, positionId, amount)
}

func (e *Exchange) DepositAsMessage(amount *big.Int, publicKey string, positionId *big.Int) (to common.Address, payload []byte, err error) {
	tokenIdUint, _ := big.NewInt(0).SetString(strings.TrimPrefix(tokenId[e.NetworkId], "0x"), 16)
	encoder, _ := abi.ExchangeMetaData.GetAbi()
	pack, err := encoder.Pack("deposit0", e.PublicKeyToUint256(publicKey), tokenIdUint, positionId, amount)

	return common.HexToAddress(exchangeAddress[e.NetworkId]), pack, err
}

func (e *Exchange) GetWithdrawalBalance(publicKey string) (*big.Int, error) {
	tokenIdUint, _ := big.NewInt(0).SetString(strings.TrimPrefix(tokenId[e.NetworkId], "0x"), 16)
	return e.exchange.GetWithdrawalBalance(&bind.CallOpts{}, e.PublicKeyToUint256(publicKey), tokenIdUint)
}

func (e *Exchange) Withdraw(publicKey string, transact *bind.TransactOpts) (*types.Transaction, error) {
	if transact.Signer == nil {
		transact.Signer = e.ethSigner
	}
	tokenIdUint, _ := big.NewInt(0).SetString(strings.TrimPrefix(tokenId[e.NetworkId], "0x"), 16)
	return e.exchange.Withdraw(transact, e.PublicKeyToUint256(publicKey), tokenIdUint)
}

func (e *Exchange) WithdrawAsMessage(publicKey string) (to common.Address, payload []byte, err error) {
	tokenIdUint, _ := big.NewInt(0).SetString(strings.TrimPrefix(tokenId[e.NetworkId], "0x"), 16)
	encoder, _ := abi.ExchangeMetaData.GetAbi()
	pack, err := encoder.Pack("withdraw", e.PublicKeyToUint256(publicKey), tokenIdUint)
	return common.HexToAddress(exchangeAddress[e.NetworkId]), pack, err
}

func (e *Exchange) RegisterUser(signature []byte, publicKey string, transact *bind.TransactOpts) (*types.Transaction, error) {
	if transact.Signer == nil {
		transact.Signer = e.ethSigner
	}
	return e.exchange.RegisterUser(transact, e.Address, e.PublicKeyToUint256(publicKey), signature)
}

func (e *Exchange) RegisterUserAsMessage(signature []byte, publicKey string) (to common.Address, payload []byte, err error) {
	encoder, _ := abi.ExchangeMetaData.GetAbi()
	pack, err := encoder.Pack("registerUser", e.Address, e.PublicKeyToUint256(publicKey), signature)
	return common.HexToAddress(exchangeAddress[e.NetworkId]), pack, err
}
