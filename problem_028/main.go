// Answer: 669171001
// Problem 28: Number Spiral Diagonals
// What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral?

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func solve() int64 {
	// For a spiral of size n x n (n odd), the corner values at each layer are:
	// Sum of corners at layer n: 4n^2 - 6(n-1)
	sum := int64(1) // Center

	for n := int64(3); n <= 1001; n += 2 {
		sum += 4*n*n - 6*(n-1)
	}

	return sum
}

func main() { bench.Run(28, solve) }
