// Problem 169: Exploring the number of different ways a number can be expressed
// as a sum of powers of 2
// Answer: 178653872807

package main

import (
	"math/big"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	// 10^25 in binary using big.Int
	n := new(big.Int).Exp(big.NewInt(10), big.NewInt(25), nil)

	// Extract binary digits MSB first
	nbits := n.BitLen()
	bits := make([]int, nbits)
	for i := 0; i < nbits; i++ {
		bits[nbits-1-i] = int(n.Bit(i))
	}

	// Process MSB first: track (f(n_partial), f(n_partial - 1))
	fa := int64(1) // f(1)
	fb := int64(1) // f(0)

	for i := 1; i < len(bits); i++ {
		if bits[i] == 0 {
			// n_partial = 2 * n_partial
			// f(2k) = f(k) + f(k-1), f(2k-1) = f(k-1)
			newFa := fa + fb
			newFb := fb
			fa, fb = newFa, newFb
		} else {
			// n_partial = 2 * n_partial + 1
			// f(2k+1) = f(k), f(2k) = f(k) + f(k-1)
			newFa := fa
			newFb := fa + fb
			fa, fb = newFa, newFb
		}
	}
	return fa
}

func main() { bench.Run(169, solve) }
