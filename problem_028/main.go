// Problem 28: Number Spiral Diagonals
// What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral?
// Answer: 669171001

package main

import (
	"fmt"
	"time"
)

func solve() int64 {
	// For a spiral of size n x n (n odd), the corner values at each layer are:
	// Sum of corners at layer n: 4n^2 - 6(n-1)
	sum := int64(1) // Center

	for n := int64(3); n <= 1001; n += 2 {
		sum += 4*n*n - 6*(n-1)
	}

	return sum
}

func benchmark(iterations int) time.Duration {
	// Warmup
	for i := 0; i < 10; i++ {
		solve()
	}

	start := time.Now()
	var result int64
	for i := 0; i < iterations; i++ {
		result = solve()
	}
	elapsed := time.Since(start)
	fmt.Printf("Result: %d (%.2f ns/op)\n", result, float64(elapsed.Nanoseconds())/float64(iterations))
	return elapsed
}

func main() {
	const iterations = 10000

	fmt.Println("Problem 28: Number Spiral Diagonals")
	fmt.Println("====================================")
	fmt.Printf("Sum of diagonals in 1001x1001 spiral, Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
