// Problem 9: Special Pythagorean Triplet
// Find the Pythagorean triplet where a + b + c = 1000.

package main

import (
	"fmt"
	"time"
)

const target = 1000

// Naive: O(n²) - double loop, derive c from sum constraint
func naive(s int) int {
	for a := 1; a < s; a++ {
		for b := a + 1; b < s; b++ {
			c := s - a - b
			if c <= b {
				break // c must be > b
			}
			if a*a+b*b == c*c {
				return a * b * c
			}
		}
	}
	return 0
}

// Optimized: O(n) - derive b algebraically from both constraints
// Given: a + b + c = s and a² + b² = c²
// Substituting c = s - a - b into a² + b² = c²:
// a² + b² = (s - a - b)²
// a² + b² = s² - 2sa - 2sb + a² + 2ab + b²
// 0 = s² - 2sa - 2sb + 2ab
// b(2s - 2a) = s² - 2sa
// b = (s² - 2sa) / (2s - 2a)
func optimized(s int) int {
	for a := 1; a < s/3; a++ {
		numerator := s*s - 2*s*a
		denominator := 2*s - 2*a
		if numerator%denominator == 0 {
			b := numerator / denominator
			c := s - a - b
			if a < b && b < c && a*a+b*b == c*c {
				return a * b * c
			}
		}
	}
	return 0
}

func benchmark(name string, f func(int) int, s int, iterations int) time.Duration {
	// Warmup
	for i := 0; i < 10; i++ {
		f(s)
	}

	start := time.Now()
	for i := 0; i < iterations; i++ {
		f(s)
	}
	elapsed := time.Since(start)
	result := f(s)
	fmt.Printf("%s: %d (%.2f ns/op)\n", name, result, float64(elapsed.Nanoseconds())/float64(iterations))
	return elapsed
}

func main() {
	const iterations = 10000

	fmt.Println("Problem 9: Special Pythagorean Triplet")
	fmt.Println("======================================")
	fmt.Printf("Target sum: %d, Iterations: %d\n\n", target, iterations)

	naiveTime := benchmark("Naive    ", naive, target, iterations)
	optTime := benchmark("Optimized", optimized, target, iterations)

	fmt.Printf("\nSpeedup: %.2fx\n", float64(naiveTime)/float64(optTime))
}
