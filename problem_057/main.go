// Problem 57: Square Root Convergents
// In the first one-thousand expansions of the continued fraction for sqrt(2),
// how many fractions contain a numerator with more digits than the denominator?
// Answer: 153
//
// The continued fraction expansion of sqrt(2) = 1 + 1/(2 + 1/(2 + 1/(2 + ...)))
// Starting: n/d = 1/1, then iteratively: n' = n + 2*d, d' = n + d
// (because 1 + 1/(1 + n/d) = 1 + d/(d+n) = (2d+n)/(d+n))

package main

import (
	"math/big"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	count := 0
	n := big.NewInt(1)
	d := big.NewInt(1)
	tmp := new(big.Int)

	for i := 0; i < 1000; i++ {
		// n' = n + 2*d, d' = n + d
		newN := new(big.Int).Add(n, tmp.Mul(big.NewInt(2), d))
		newD := new(big.Int).Add(n, d)
		n = newN
		d = newD

		if len(n.String()) > len(d.String()) {
			count++
		}
	}
	return int64(count)
}

func main() { bench.Run(57, solve) }
