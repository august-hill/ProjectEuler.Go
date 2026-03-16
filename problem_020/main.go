// Problem 20: Factorial Digit Sum
// Find the sum of the digits in 100!
// Answer: 648

package main

import (
	"fmt"
	"math/big"
	"time"
)

func solve() int {
	factorial := big.NewInt(1)
	for n := 2; n <= 100; n++ {
		factorial.Mul(factorial, big.NewInt(int64(n)))
	}

	sum := 0
	for _, ch := range factorial.String() {
		sum += int(ch - '0')
	}
	return sum
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
	const iterations = 10000

	fmt.Println("Problem 20: Factorial Digit Sum")
	fmt.Println("================================")
	fmt.Printf("Computing sum of digits in 100!, Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
