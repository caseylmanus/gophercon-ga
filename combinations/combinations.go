package combinations

import "math/big"

func Possible(length, possibilities int64) *big.Int {
	n := big.NewInt(possibilities)
	k := big.NewInt(length)

	// Calculate n!
	nFactorial := new(big.Int).MulRange(1, n.Int64())

	// Calculate k!
	kFactorial := new(big.Int).MulRange(1, k.Int64())

	// Calculate (n - k)!
	nMinusK := new(big.Int).Sub(n, k)
	nMinusKFactorial := new(big.Int).MulRange(1, nMinusK.Int64())

	// Calculate C(n, k)
	combinations := new(big.Int).Div(nFactorial, new(big.Int).Mul(kFactorial, nMinusKFactorial))

	return combinations
}
