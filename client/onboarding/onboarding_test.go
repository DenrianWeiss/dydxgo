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
		signer: signerFn,
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

func TestOnBoarding_DeriveStarkKey(t *testing.T) {
	// todo
}
