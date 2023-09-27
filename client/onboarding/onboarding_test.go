package onboarding

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"github.com/denrianweiss/dydxgo/client/base"
	"github.com/denrianweiss/dydxgo/constants"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"log"
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
			NetworkId: constants.NetworkIdGoerli,
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
			NetworkId: constants.NetworkIdGoerli,
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
