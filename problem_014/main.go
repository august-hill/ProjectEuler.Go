// Problem 14: Longest Collatz Sequence
// Find starting number under 1 million with longest Collatz chain.

package main

import (
	"fmt"
	"time"
)

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

func solve() uint64 {
	var bestN uint64 = 0
	bestLen := 0

	for n := uint64(1); n < limit; n++ {
		length := collatzLength(n)
		if length > bestLen {
			bestLen = length
			bestN = n
		}
	}
	return bestN
}

func benchmark(iterations int) time.Duration {
	// Warmup
	for i := 0; i < 10; i++ {
		solve()
	}

	start := time.Now()
	var result uint64
	for i := 0; i < iterations; i++ {
		result = solve()
	}
	elapsed := time.Since(start)
	fmt.Printf("Result: %d (%.2f ms/op)\n", result, float64(elapsed.Milliseconds())/float64(iterations))
	return elapsed
}

func main() {
	const iterations = 10

	fmt.Println("Problem 14: Longest Collatz Sequence")
	fmt.Println("=====================================")
	fmt.Printf("Limit: %d, Iterations: %d\n\n", limit, iterations)

	benchmark(iterations)
}
