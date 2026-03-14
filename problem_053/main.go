// Problem 53: Combinatoric Selections
// Count C(n,r) > 1,000,000 for 1 <= n <= 100.
// Answer: 4075

package main

import (
	"fmt"
	"time"
)

func solve() int {
	// Use Pascal's triangle to avoid overflow
	// C(n,r) = C(n-1,r-1) + C(n-1,r)
	const limit = 1000000

	count := 0
	prev := make([]int, 102)
	prev[0] = 1

	for n := 1; n <= 100; n++ {
		curr := make([]int, 102)
		curr[0] = 1
		for r := 1; r <= n; r++ {
			curr[r] = prev[r-1] + prev[r]
			// Cap at limit+1 to detect > limit without overflow
			if curr[r] > limit {
				curr[r] = limit + 1
				count++
			}
		}
		prev = curr
	}

	return count
}

func benchmark(iterations int) time.Duration {
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
	fmt.Println("Problem 53: Combinatoric Selections")
	fmt.Println("====================================")
	benchmark(100000)
}
