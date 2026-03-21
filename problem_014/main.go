// Answer: 837799
// Problem 14: Longest Collatz Sequence
// Find starting number under 1 million with longest Collatz chain.

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const limit = 1_000_000

func collatzLength(n uint64) int {
	length := 1
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		length++
	}
	return length
}

func solve() int64 {
	var bestN uint64 = 0
	bestLen := 0

	for n := uint64(1); n < limit; n++ {
		length := collatzLength(n)
		if length > bestLen {
			bestLen = length
			bestN = n
		}
	}
	return int64(bestN)
}

func main() { bench.Run(14, solve) }
