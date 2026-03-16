// Problem 21: Amicable Numbers
// Evaluate the sum of all amicable numbers under 10000.
// Answer: 31626

package main

import (
	"fmt"
	"math"
	"time"
)

func sumProperDivisors(n int) int {
	if n <= 1 {
		return 0
	}
	sum := 1
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			sum += i
			other := n / i
			if other != i {
				sum += other
			}
		}
	}
	return sum
}

func solve() int {
	sum := 0
	for a := 2; a < 10000; a++ {
		b := sumProperDivisors(a)
		if b != a && b < 10000 && sumProperDivisors(b) == a {
			sum += a
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
	const iterations = 1000

	fmt.Println("Problem 21: Amicable Numbers")
	fmt.Println("=============================")
	fmt.Printf("Sum of amicable numbers under 10000, Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
