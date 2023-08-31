package ethereum

import (
	"github.com/denrianweiss/dydxgo/client/base"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ETHClient struct {
	base.BaseUser
	ethSigner bind.SignerFn
	ethClient *ethclient.Client
}

type ETH struct {
	Exchange *Exchange
	Token    *Token
}

func (e *ETH) New(user base.BaseUser, signer bind.SignerFn, client *ethclient.Client) *ETH {
	ec := ETHClient{BaseUser: user, ethClient: client, ethSigner: signer}
	e.Exchange = (&Exchange{}).New(ec)
	e.Token = (&Token{}).New(ec)
	return e
}
