// Problem 33: Digit Cancelling Fractions
// Find the denominator of the product of the four "curious" fractions in lowest terms.
// Answer: 100

package main

import (
	"fmt"
	"time"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func solve() int {
	numProduct := 1
	denProduct := 1

	for a := 1; a <= 9; a++ {
		for b := 1; b <= 9; b++ {
			for c := 1; c <= 9; c++ {
				for d := 1; d <= 9; d++ {
					num := a*10 + b
					den := c*10 + d

					// Fraction must be less than 1
					if num >= den {
						continue
					}

					// Check if "cancelling" gives equivalent fraction
					// ab/cd: if b == c, check if a/d == (ab)/(cd)
					if b == c && a*den == num*d {
						numProduct *= num
						denProduct *= den
					}
				}
			}
		}
	}

	return denProduct / gcd(numProduct, denProduct)
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

	fmt.Println("Problem 33: Digit Cancelling Fractions")
	fmt.Println("=======================================")
	fmt.Printf("Denominator of product in lowest terms, Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
