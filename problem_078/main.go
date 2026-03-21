// Problem 78: Coin Partitions
// Find the least value of n for which p(n) is divisible by one million.
// Answer: 55374

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

const (
	pLimit078  = 100000
	modulus078 = 1000000
)

func solve() int64 {
	p := make([]int, pLimit078)
	p[0] = 1

	// Precompute generalized pentagonal numbers
	pentagonals := make([]int, 0, 1000)
	for k := 1; k < 500; k++ {
		pentagonals = append(pentagonals, k*(3*k-1)/2)
		pentagonals = append(pentagonals, k*(3*k+1)/2)
	}

	for n := 1; n < pLimit078; n++ {
		sum := int64(0)
		for i, pent := range pentagonals {
			if pent > n {
				break
			}

			sign := int64(1)
			if i%4 >= 2 {
				sign = -1
			}
			sum = (sum + sign*int64(p[n-pent])) % modulus078
		}

		p[n] = int(((sum % modulus078) + modulus078) % modulus078)

		if p[n] == 0 {
			return int64(n)
		}
	}

	return 0
}

func main() { bench.Run(78, solve) }
