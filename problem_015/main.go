// Problem 15: Lattice Paths
// Count routes through 20x20 grid (only right/down moves).
// Answer: C(40, 20) = 137846528820

package main

import (
	"fmt"
	"time"
)

// C(2n, n) computed without overflow by interleaving multiply/divide
func latticePaths(n uint64) uint64 {
	result := uint64(1)
	for i := uint64(1); i <= n; i++ {
		result = result * (n + i) / i
	}
	return result
}

func solve() uint64 {
	return latticePaths(20)
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
	fmt.Printf("Result: %d (%.2f ns/op)\n", result, float64(elapsed.Nanoseconds())/float64(iterations))
	return elapsed
}

func main() {
	const iterations = 100000

	fmt.Println("Problem 15: Lattice Paths")
	fmt.Println("=========================")
	fmt.Printf("Grid: 20x20, Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
