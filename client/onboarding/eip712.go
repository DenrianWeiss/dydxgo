package onboarding

import (
	"github.com/denrianweiss/dydxgo/constants"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"log"
	"math/big"
)

func (b *OnBoarding) CreateMessage(msg map[string]interface{}) apitypes.TypedData {
	d := apitypes.TypedData{
		Types: map[string][]apitypes.Type{
			"EIP712Domain": []apitypes.Type{
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "version",
					Type: "string",
				},
				{
					Name: "chainId",
					Type: "uint256",
				},
			},
		},
		Domain: apitypes.TypedDataDomain{
			Name:              constants.Domain,
			Version:           constants.Version,
			ChainId:           (*math.HexOrDecimal256)(big.NewInt(b.NetworkId)),
			VerifyingContract: "",
			Salt:              "",
		},
		PrimaryType: constants.Domain,
		Message:     msg,
	}
	d.Types[constants.Domain] = []apitypes.Type{
		{
			Name: "action",
			Type: "string",
		},
	}
	if b.NetworkId == constants.NetworkIdMainnet {
		msg["onlySignOn"] = constants.OnlySignOnDomainMainnet
		d.Types[constants.Domain] = []apitypes.Type{
			{
				Name: "action",
				Type: "string",
			},
			{
				Name: "onlySignOn",
				Type: "string",
			},
		}
	}
	return d
}

func (b *OnBoarding) HashMessage(d apitypes.TypedData) ([]byte, error) {
	sig, _, err := apitypes.TypedDataAndHash(d)
	if err != nil {
		log.Panic("failed to encode typed data.")
	}
	return sig, err
}

func (b *OnBoarding) SignMessage(d apitypes.TypedData) ([]byte, error) {
	sig, msg, err := apitypes.TypedDataAndHash(d)
	if err != nil {
		log.Panic("failed to encode typed data.")
	}
	log.Printf("Signing Typed Data %x", msg)
	res, err := b.CryptoSigner(sig)
	// Fix Signature
	lastByte := res[len(res)-1]
	if lastByte <= 2 {
		res[len(res)-1] = lastByte + 27
	}
	res = append(res, 0)
	return res, err
}
