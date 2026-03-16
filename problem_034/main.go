// Problem 34: Digit Factorials
// Find the sum of all numbers which are equal to the sum of the factorial of their digits.
// Answer: 40730

package main

import (
	"fmt"
	"time"
)

var factorials = [10]int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

func digitFactorialSum(n int) int {
	sum := 0
	for n > 0 {
		sum += factorials[n%10]
		n /= 10
	}
	return sum
}

func solve() int {
	// Upper bound: 7 * 9! = 2540160
	sum := 0
	for n := 3; n <= 2540160; n++ {
		if n == digitFactorialSum(n) {
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
	const iterations = 10

	fmt.Println("Problem 34: Digit Factorials")
	fmt.Println("=============================")
	fmt.Printf("Sum of numbers equal to sum of digit factorials, Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
