package ec

import (
	"github.com/consensys/gnark-crypto/ecc/stark-curve"
	"math/big"
)

func DerivePublicKey(private *big.Int) *starkcurve.G1Affine {
	affine := starkcurve.G1Affine{}
	coords := affine.ScalarMultiplicationBase(private)
	return coords
}
