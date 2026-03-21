// Problem 70: Totient Permutation
// Find n, 1 < n < 10^7, for which phi(n) is a permutation of n and n/phi(n) is minimized.
// Answer: 8319823

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

const limit070 = 10000000

func isPermutation(a, b int) bool {
	var digitsA, digitsB [10]int
	for a > 0 {
		digitsA[a%10]++
		a /= 10
	}
	for b > 0 {
		digitsB[b%10]++
		b /= 10
	}
	return digitsA == digitsB
}

func solve() int64 {
	phi := make([]int, limit070)
	for i := range phi {
		phi[i] = i
	}

	for i := 2; i < limit070; i++ {
		if phi[i] == i {
			for j := i; j < limit070; j += i {
				phi[j] -= phi[j] / i
			}
		}
	}

	bestN := 0
	bestRatio := 1e18

	for n := 2; n < limit070; n++ {
		if isPermutation(n, phi[n]) {
			ratio := float64(n) / float64(phi[n])
			if ratio < bestRatio {
				bestRatio = ratio
				bestN = n
			}
		}
	}

	return int64(bestN)
}

func main() { bench.Run(70, solve) }
