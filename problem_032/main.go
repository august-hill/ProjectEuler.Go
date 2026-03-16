// Problem 32: Pandigital Products
// Find the sum of all products whose multiplicand/multiplier/product identity
// can be written as a 1 through 9 pandigital.
// Answer: 45228

package main

import (
	"fmt"
	"time"
)

func isPandigital(a, b, c int) bool {
	s := fmt.Sprintf("%d%d%d", a, b, c)
	if len(s) != 9 {
		return false
	}

	var digits [10]bool
	for _, ch := range s {
		d := int(ch - '0')
		if d == 0 || digits[d] {
			return false
		}
		digits[d] = true
	}
	return true
}

func solve() int {
	products := make(map[int]bool)

	for a := 1; a < 100; a++ {
		start := 1000
		end := 9999
		if a >= 10 {
			start = 100
			end = 999
		}

		for b := start; b <= end; b++ {
			c := a * b
			if isPandigital(a, b, c) {
				products[c] = true
			}
		}
	}

	sum := 0
	for p := range products {
		sum += p
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

	fmt.Println("Problem 32: Pandigital Products")
	fmt.Println("================================")
	fmt.Printf("Sum of pandigital products, Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
