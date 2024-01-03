package onboarding

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/denrianweiss/dydxgo/ec"
	"math/big"
)

func (b *OnBoarding) CreateAccount(ethereumAddress string) (*OnboardingResponse, string, error) {
	// Stark key
	privateKey := b.DeriveStarkKey(ethereumAddress)
	pkInt, _ := new(big.Int).SetString(privateKey, 0)
	publicKey := ec.DerivePublicKey(pkInt)
	// Convert public key to what dydx's api expect.
	pkx := fmt.Sprintf("%x", publicKey.X.Bytes())
	pky := fmt.Sprintf("%x", publicKey.Y.Bytes())
	data := map[string]string{
		"starkKey":            pkx,
		"starkKeyYCoordinate": pky,
	}
	apiKeySignature := b.GetDefaultApiCredentialsSignature()
	res, err := b.post("onboarding", data, ethereumAddress, apiKeySignature)
	if err != nil {
		return nil, "", err
	}
	result := &OnboardingResponse{}
	if err = json.Unmarshal(res, result); err != nil {
		return nil, "", errors.New("json parser error")
	}
	return result, privateKey, nil
}
