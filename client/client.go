package client

import (
	"github.com/denrianweiss/dydxgo/client/base"
	"github.com/denrianweiss/dydxgo/client/ethereum"
	"github.com/denrianweiss/dydxgo/client/onboarding"
	"github.com/denrianweiss/dydxgo/client/private"
	"github.com/denrianweiss/dydxgo/client/public"
	"github.com/denrianweiss/dydxgo/signer"
	"github.com/denrianweiss/dydxgo/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"net/http"
	"strings"
)

type Options struct {
	NetworkId              int
	Host                   string
	DefaultEthereumAddress string

	StarkPrivateKey   string
	ApiKeyCredentials *types.ApiKeyCredentials

	HttpClient   *http.Client
	Logger       *log.Logger
	CryptoSigner signer.CryptoSigner
	EthSigner    bind.SignerFn
	EthConn      *ethclient.Client
}

type Client struct {
	base.BaseUser
	Eth        *ethereum.ETH
	OnBoarding *onboarding.OnBoarding
	Private    *private.Private
	Public     *public.Public
}

func New(options Options) Client {
	// Fill Basic Info First.
	clientInstance := Client{}
	clientInstance.Address = common.HexToAddress(options.DefaultEthereumAddress)
	clientInstance.NetworkId = int64(options.NetworkId)
	if options.CryptoSigner != nil {
		ob := onboarding.OnBoarding{
			CryptoSigner: options.CryptoSigner,
		}
		ob.NetworkId = int64(options.NetworkId)
		clientInstance.OnBoarding = &ob
	}
	if options.StarkPrivateKey == "" {
		// Start OnBoarding Process.
		privateKey := clientInstance.OnBoarding.DeriveStarkKey(clientInstance.Address.String())
		clientInstance.StarkPrivateKey = privateKey
	} else {
		clientInstance.StarkPrivateKey = options.StarkPrivateKey
	}
	// Remove 0x Prefix for private key
	trimmed := strings.TrimPrefix(clientInstance.StarkPrivateKey, "0x")
	clientInstance.StarkPrivateKey = trimmed
	if options.ApiKeyCredentials == nil {
		apiKey := clientInstance.OnBoarding.RecoverDefaultApiCredentials(clientInstance.Address.String())
		clientInstance.ApiKeyCredentials = apiKey
	} else {
		clientInstance.ApiKeyCredentials = options.ApiKeyCredentials
	}

	// Create sub instances.
	if options.EthSigner != nil && options.EthConn != nil {
		clientInstance.Eth = (&ethereum.ETH{}).New(
			clientInstance.BaseUser,
			options.EthSigner,
			options.EthConn,
		)
	}

	clientInstance.Private = &private.Private{
		BaseUser:   clientInstance.BaseUser,
		Host:       options.Host,
		HttpClient: options.HttpClient,
		RateLimit:  &types.RateLimit{},
		Logger:     options.Logger,
	}

	clientInstance.Public = &public.Public{
		BaseUser:   clientInstance.BaseUser,
		Host:       options.Host,
		HttpClient: options.HttpClient,
		RateLimit:  &types.RateLimit{},
		Logger:     options.Logger,
	}
	return clientInstance
}
