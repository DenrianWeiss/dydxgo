package ec

import (
	"math/big"
	"os"
	"testing"
)

func TestGeneratePublicKey(t *testing.T) {
	starkPrivate := os.Getenv("STARK_PRIVATE")
	private, success := big.NewInt(0).SetString(starkPrivate, 16)
	if !success {
		t.Log("Please set stark private key properly in environ")
		t.Fail()
	}
	DerivePublicKey(private)
}
