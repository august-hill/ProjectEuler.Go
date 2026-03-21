// Problem 48: Self Powers
// Find last 10 digits of 1^1 + 2^2 + 3^3 + ... + 1000^1000.
// Answer: 9110846700

package main

import (
	"math/big"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	mod := big.NewInt(10000000000) // 10^10
	sum := big.NewInt(0)

	for i := int64(1); i <= 1000; i++ {
		base := big.NewInt(i)
		power := new(big.Int).Exp(base, base, mod)
		sum.Add(sum, power)
		sum.Mod(sum, mod)
	}
	return sum.Int64()
}

func main() { bench.Run(48, solve) }
