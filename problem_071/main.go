// Problem 71: Ordered Fractions
// Find the numerator of the fraction immediately to the left of 3/7
// in the Farey sequence with d <= 1,000,000.
// Answer: 428570

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	limit := 1000000
	bestN := 0
	bestD := 1

	for d := 2; d <= limit; d++ {
		n := (3*d - 1) / 7
		if n*bestD > bestN*d {
			bestN = n
			bestD = d
		}
	}

	_ = bestD
	return int64(bestN)
}

func main() { bench.Run(71, solve) }
