package onboarding

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/denrianweiss/dydxgo/client/base"
	"github.com/denrianweiss/dydxgo/constants"
	"github.com/denrianweiss/dydxgo/signer"
	"github.com/denrianweiss/dydxgo/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"log"
	"math/big"
	"strings"
)

type OnBoarding struct {
	base.BaseClient
	CryptoSigner signer.CryptoSigner
}

func (b *OnBoarding) DeriveStarkKey(ethereumAddress string) string {
	message := b.CreateMessage(map[string]interface{}{"action": constants.OffChainKeyDerivationAction})
	sig, err := b.SignMessage(message)
	if err != nil {
		log.Panic(err)
	}
	sha3 := solsha3.SoliditySHA3([]string{"uint256"}, fmt.Sprintf("0x%x", sig))
	hashedSignature := hexutil.Encode(sha3)

	privateKey, _ := new(big.Int).SetString(hashedSignature, 0)
	privateKey = new(big.Int).Rsh(privateKey, 5)
	return fmt.Sprintf("0x%s", privateKey.Text(16))
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
