// Problem 26: Reciprocal Cycles
// Find the value of d < 1000 for which 1/d contains the longest recurring cycle.
// Answer: 983

package main

import (
	"fmt"
	"time"
)

func cycleLength(d int) int {
	seen := make([]int, d)
	for i := range seen {
		seen[i] = -1
	}

	remainder := 1
	position := 0

	for remainder != 0 {
		if seen[remainder] >= 0 {
			return position - seen[remainder]
		}
		seen[remainder] = position
		remainder = (remainder * 10) % d
		position++
	}

	return 0 // Terminating decimal
}

func solve() int {
	maxCycle := 0
	result := 0

	for d := 2; d < 1000; d++ {
		cycle := cycleLength(d)
		if cycle > maxCycle {
			maxCycle = cycle
			result = d
		}
	}

	return result
}

func benchmark(iterations int) time.Duration {
	// Warmup
	for i := 0; i < 10; i++ {
		solve()
	}

	start := time.Now()
	var result int
	for i := 0; i < iterations; i++ {
		result = solve()
	}
	elapsed := time.Since(start)
	fmt.Printf("Result: %d (%.2f ns/op)\n", result, float64(elapsed.Nanoseconds())/float64(iterations))
	return elapsed
}

func main() {
	const iterations = 1000

	fmt.Println("Problem 26: Reciprocal Cycles")
	fmt.Println("==============================")
	fmt.Printf("Finding d < 1000 with longest recurring cycle in 1/d, Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
