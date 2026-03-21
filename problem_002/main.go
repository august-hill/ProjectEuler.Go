// Answer: 4613732
// Problem 002: Even Fibonacci Numbers
// Find the sum of all even-valued Fibonacci terms below 4 million.

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const limit = 4_000_000

// evenRecurrence uses E(n) = 4*E(n-1) + E(n-2) to generate only even Fibs
func solve() int64 {
	sum := 0
	a, b := 2, 8
	for a < limit {
		sum += a
		a, b = b, 4*b+a
	}
	return int64(sum)
}

func main() { bench.Run(2, solve) }
