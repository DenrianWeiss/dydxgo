package onboarding

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/denrianweiss/dydxgo/client/base"
	"github.com/denrianweiss/dydxgo/constants"
	"github.com/denrianweiss/dydxgo/ec"
	"github.com/denrianweiss/dydxgo/signer"
	"github.com/denrianweiss/dydxgo/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"io"
	"log"
	"math/big"
	"net/http"
	"strings"
)

type OnBoarding struct {
	base.BaseClient
	Host         string
	CryptoSigner signer.CryptoSigner
}

func (b *OnBoarding) DeriveStarkKey(ethereumAddress string) string {
	message := b.CreateMessage(map[string]interface{}{"action": constants.OffChainKeyDerivationAction})
	sig, err := b.SignMessage(message)
	if err != nil {
		log.Panic(err)
	}
	sha3 := solsha3.SoliditySHA3(fmt.Sprintf("0x%x", sig))
	hashedSignature := hexutil.Encode(sha3)

	privateKey, _ := new(big.Int).SetString(hashedSignature, 0)
	privateKey = new(big.Int).Rsh(privateKey, 5)
	return fmt.Sprintf("0x%s", privateKey.Text(16))
}

func (b *OnBoarding) GetDefaultApiCredentialsSignature() string {
	message := b.CreateMessage(map[string]interface{}{"action": constants.OffChainOnboardingAction})
	sig, err := b.SignMessage(message)
	if err != nil {
		log.Panic(err)
	}
	signature := fmt.Sprintf("0x%x", sig)
	return signature
}

func (b *OnBoarding) RecoverDefaultApiCredentials(ethereumAddress string) *types.ApiKeyCredentials {
	message := b.CreateMessage(map[string]interface{}{"action": constants.OffChainOnboardingAction})
	sig, err := b.SignMessage(message)
	if err != nil {
		log.Panic(err)
	}
	signature := fmt.Sprintf("0x%x", sig)
	rHex := signature[2:66]
	rInt, _ := big.NewInt(0).SetString(rHex, 16)

	hashedRBytes := solsha3.SoliditySHA3([]string{"uint256"}, rInt.String())
	secretBytes := hashedRBytes[:30]
	sHex := signature[66:130]
	sInt, _ := big.NewInt(0).SetString(sHex, 16)

	hashedSBytes := solsha3.SoliditySHA3([]string{"uint256"}, sInt.String())
	keyBytes := hashedSBytes[:16]
	passphraseBytes := hashedSBytes[16:31]

	keyHex := hex.EncodeToString(keyBytes)
	keyUuid := strings.Join([]string{
		keyHex[:8],
		keyHex[8:12],
		keyHex[12:16],
		keyHex[16:20],
		keyHex[20:],
	}, "-")
	return &types.ApiKeyCredentials{
		Secret:     base64.URLEncoding.EncodeToString(secretBytes),
		Key:        keyUuid,
		Passphrase: base64.URLEncoding.EncodeToString(passphraseBytes),
	}
}

func (b *OnBoarding) post(endpoint string, data interface{}, ethereumAddress, onboardingSig string) (resp []byte, err error) {
	// Marshal data to json
	url := fmt.Sprintf("%s/v3/%s", b.Host, endpoint)
	// Create request
	dataJson, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dataJson))
	if err != nil {
		log.Panic(err)
	}
	if onboardingSig == "" {
		onboardingSig = b.GetDefaultApiCredentialsSignature()
	}
	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("DYDX-SIGNATURE", onboardingSig)
	req.Header.Set("DYDX-ETHEREUM-ADDRESS", ethereumAddress)
	// Send request
	client := http.DefaultClient
	respVal, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if respVal.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", respVal.StatusCode)
	}
	defer respVal.Body.Close()
	// Read response
	resp, err = io.ReadAll(respVal.Body)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *OnBoarding) Onboard(ethereumAddress string, onboardingParams map[string]interface{}) {
	if onboardingParams == nil {
		starkKey := b.DeriveStarkKey(ethereumAddress)
		pkInt, _ := new(big.Int).SetString(starkKey, 0)
		publicKey := ec.DerivePublicKey(pkInt)
		pkx := fmt.Sprintf("%x", publicKey.X.Bytes())
		pky := fmt.Sprintf("%x", publicKey.Y.Bytes())
		onboardingParams = map[string]interface{}{
			"starkKey":            pkx,
			"starkKeyYCoordinate": pky,
		}
	}
	b.post("onboarding", onboardingParams, ethereumAddress, "")
}
