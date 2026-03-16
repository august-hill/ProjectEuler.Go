// Problem 25: 1000-digit Fibonacci Number
// What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
// Answer: 4782

package main

import (
	"fmt"
	"math/big"
	"time"
)

func solve() int {
	fib1 := big.NewInt(1) // F(1)
	fib2 := big.NewInt(1) // F(2)
	term := 2

	// 10^999 is the smallest 1000-digit number
	threshold := new(big.Int).Exp(big.NewInt(10), big.NewInt(999), nil)

	for fib2.Cmp(threshold) < 0 {
		fib1.Add(fib1, fib2)
		fib1, fib2 = fib2, fib1
		term++
	}

	return term
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
	const iterations = 100

	fmt.Println("Problem 25: 1000-digit Fibonacci Number")
	fmt.Println("========================================")
	fmt.Printf("Finding index of first 1000-digit Fibonacci, Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
