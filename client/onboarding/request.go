package onboarding

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/denrianweiss/dydxgo/ec"
	"io"
	"log"
	"math/big"
	"net/http"
	"strings"
)

func (b *OnBoarding) post(endpoint string, data map[string]string, ethereumAddress string, signature string) ([]byte, error) {
	// Json serialize data
	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	requestPath := fmt.Sprintf("%s%s", b.Host, fmt.Sprintf("/v3/%s", endpoint))
	req, err := http.NewRequest(http.MethodPost, requestPath, strings.NewReader(string(serialized)))
	req.Header.Add("DYDX-SIGNATURE", signature)
	req.Header.Add("DYDX-ETHEREUM-ADDRESS", ethereumAddress)
	req.Header.Add("content-type", "application/json")
	result := &OnboardingResponse{}
	resRaw, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	res, _ := io.ReadAll(resRaw.Body)
	// Log body
	log.Printf("uri: %s, response body: %s", requestPath, res)
	if err = json.Unmarshal(res, result); err != nil {
		return nil, errors.New("json parser error")
	}
	return res, nil
}

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
