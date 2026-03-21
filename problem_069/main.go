// Problem 069: Totient Maximum
// Find n <= 1,000,000 for which n/phi(n) is a maximum.
// Answer: 510510
//
// n/phi(n) is maximized when n is the product of consecutive primes.
// 2 * 3 * 5 * 7 * 11 * 13 * 17 = 510510

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	limit := 1000000
	result := 1

	for _, p := range primes {
		if result*p > limit {
			break
		}
		result *= p
	}

	return int64(result)
}

func main() { bench.Run(69, solve) }
