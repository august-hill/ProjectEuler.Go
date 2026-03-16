// Problem 30: Digit Fifth Powers
// Find the sum of all numbers that can be written as the sum of fifth powers of their digits.
// Answer: 443839

package main

import (
	"fmt"
	"time"
)

var pow5 = [10]int{0, 1, 32, 243, 1024, 3125, 7776, 16807, 32768, 59049}

func fifthPowerSum(n int) int {
	sum := 0
	for n > 0 {
		sum += pow5[n%10]
		n /= 10
	}
	return sum
}

func solve() int {
	// Upper bound: 6 * 9^5 = 354294
	sum := 0
	for n := 2; n <= 354294; n++ {
		if n == fifthPowerSum(n) {
			sum += n
		}
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
	const iterations = 100

	fmt.Println("Problem 30: Digit Fifth Powers")
	fmt.Println("===============================")
	fmt.Printf("Sum of numbers equal to sum of fifth powers of digits, Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
