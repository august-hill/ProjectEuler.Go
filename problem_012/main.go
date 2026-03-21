// Answer: 76576500
// Problem 12: Highly Divisible Triangular Number
// Find the first triangle number with over 500 divisors.

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func countDivisors(n uint64) int {
	count := 0
	sqrtN := uint64(math.Sqrt(float64(n)))

	for i := uint64(1); i <= sqrtN; i++ {
		if n%i == 0 {
			if i*i == n {
				count++
			} else {
				count += 2
			}
		}
	}
	return count
}

func solve() int64 {
	n := uint64(1)
	for {
		triangle := n * (n + 1) / 2
		if countDivisors(triangle) > 500 {
			return int64(triangle)
		}
		n++
	}
}

func main() { bench.Run(12, solve) }
