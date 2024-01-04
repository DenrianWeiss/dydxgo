package onboarding

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"github.com/denrianweiss/dydxgo/client/base"
	"github.com/denrianweiss/dydxgo/constants"
	"github.com/denrianweiss/dydxgo/ec"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"log"
	"math/big"
	"os"
	"testing"
)

func TestOnBoarding_CreateMessage(t *testing.T) {

}

func TestOnBoarding_SignMessage(t *testing.T) {
	key, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	signerFn := func(sig []byte) ([]byte, error) {
		return crypto.Sign(sig, key)
	}
	// Create OnBoarding Instance
	ob := OnBoarding{
		BaseClient: base.BaseClient{
			NetworkId: constants.NetworkIdSepolia,
		},
		CryptoSigner: signerFn,
	}
	msg := ob.CreateMessage(map[string]interface{}{"action": constants.OffChainKeyDerivationAction})
	hash, err := ob.HashMessage(msg)
	log.Printf("%x\n", hash)
	if fmt.Sprintf("%x\n", hash) != "ae1353e14f103174748e254577d8bb84eddbce5477bce5dd045c9625ed49deeb" {
		t.Fail()
	}
	sig, err := ob.SignMessage(msg)
	log.Printf("sig: %x", sig)
	pub := secp256k1.CompressPubkey(key.X, key.Y)
	checked := crypto.VerifySignature(pub, hash, sig)
	if !checked {
		t.Fail()
	}
}

func TestOnBoarding_CreateAccount(t *testing.T) {
	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey == "" {
		log.Panic("PRIVATE_KEY env variable is not set")
	}
	// Derive key using private key
	key, err := crypto.HexToECDSA(privateKey)
	// Convert key to ethereum address
	addr := crypto.PubkeyToAddress(key.PublicKey)
	if err != nil {
		log.Panic(err)
	}
	signerFn := func(sig []byte) ([]byte, error) {
		return crypto.Sign(sig, key)
	}
	// Create OnBoarding Instance
	ob := OnBoarding{
		BaseClient: base.BaseClient{
			NetworkId: constants.NetworkIdSepolia,
		},
		CryptoSigner: signerFn,
		Host:         "https://api.stage.dydx.exchange",
	}
	// Create Account
	account, private, err := ob.CreateAccount(addr.String())
	if err != nil {
		log.Panic(err)
	}
	log.Printf("private: %s", private)
	log.Printf("account: %v", account)
}

func TestOnBoarding_Onboard(t *testing.T) {
	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey == "" {
		log.Panic("PRIVATE_KEY env variable is not set")
	}
	// Derive key using private key
	key, err := crypto.HexToECDSA(privateKey)
	// Convert key to ethereum address
	addr := crypto.PubkeyToAddress(key.PublicKey)
	if err != nil {
		log.Panic(err)
	}
	signerFn := func(sig []byte) ([]byte, error) {
		return crypto.Sign(sig, key)
	}
	// Create OnBoarding Instance
	ob := OnBoarding{
		BaseClient: base.BaseClient{
			NetworkId: constants.NetworkIdSepolia,
		},
		CryptoSigner: signerFn,
		Host:         "https://api.stage.dydx.exchange",
	}
	// Create Account
	ob.Onboard(addr.String(), nil)
}

func TestOnBoarding_OnBoardWithEnv(t *testing.T) {
	derivedKey := os.Getenv("DERIVED_KEY")
	if derivedKey == "" {
		log.Panic("DERIVED_KEY env variable is not set")
	}
	ethAddress := os.Getenv("ETH_ADDRESS")
	if ethAddress == "" {
		log.Panic("ETH_ADDRESS env variable is not set")
	}
	signature := os.Getenv("SIGNATURE")
	if signature == "" {
		log.Panic("SIGNATURE env variable is not set")
	}
	// Create OnBoarding Instance
	ob := OnBoarding{
		BaseClient: base.BaseClient{
			NetworkId: constants.NetworkIdSepolia,
		},
		Host:         "https://api.stage.dydx.exchange",
		CryptoSigner: nil,
	}
	pkInt, _ := new(big.Int).SetString(derivedKey, 0)
	publicKey := ec.DerivePublicKey(pkInt)
	pkx := fmt.Sprintf("%x", publicKey.X.Bytes())
	pky := fmt.Sprintf("%x", publicKey.Y.Bytes())
	onboardingParams := map[string]interface{}{
		"starkKey":            pkx,
		"starkKeyYCoordinate": pky,
	}
	// Create Account
	post, err := ob.post("onboarding", onboardingParams, ethAddress, signature+"00")
	if err != nil {
		log.Panic(err)
	}
	log.Printf("post: %s", string(post))
}
