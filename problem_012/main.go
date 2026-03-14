// Problem 12: Highly Divisible Triangular Number
// Find the first triangle number with over 500 divisors.

package main

import (
	"fmt"
	"math"
	"time"
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

func solve() uint64 {
	n := uint64(1)
	for {
		triangle := n * (n + 1) / 2
		if countDivisors(triangle) > 500 {
			return triangle
		}
		n++
	}
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

	fmt.Println("Problem 12: Highly Divisible Triangular Number")
	fmt.Println("===============================================")
	fmt.Printf("Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
