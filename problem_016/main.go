// Problem 16: Power Digit Sum
// What is the sum of the digits of 2^1000?
// Answer: 1366

package main

import (
	"fmt"
	"time"
)

// Manual digit doubling - no big integer library needed
func powerDigitSum(n int) int {
	digits := []int{1}

	for i := 0; i < n; i++ {
		carry := 0
		for j := 0; j < len(digits); j++ {
			val := digits[j]*2 + carry
			digits[j] = val % 10
			carry = val / 10
		}
		if carry > 0 {
			digits = append(digits, carry)
		}
	}

	sum := 0
	for _, d := range digits {
		sum += d
	}
	return sum
}

func solve() int {
	return powerDigitSum(1000)
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

	fmt.Println("Problem 16: Power Digit Sum")
	fmt.Println("============================")
	fmt.Printf("Computing sum of digits of 2^1000, Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
