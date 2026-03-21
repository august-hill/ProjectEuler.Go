// Answer: 25164150
// Problem 006: Sum Square Difference
// Find the difference between the square of the sum and the sum of squares
// for the first 100 natural numbers.

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

// directFormula uses the algebraic simplification:
// Difference = n(n-1)(n+1)(3n+2)/12
func solve() int64 {
	n := 100
	return int64(n * (n - 1) * (n + 1) * (3*n + 2) / 12)
}

func main() { bench.Run(6, solve) }
