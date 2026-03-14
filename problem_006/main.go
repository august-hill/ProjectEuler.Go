// Problem 006: Sum Square Difference
// Find the difference between the square of the sum and the sum of squares
// for the first 100 natural numbers.

package main

import (
	"fmt"
	"time"
)

const limit = 100

// bruteForce loops through all numbers, accumulating both sums
func bruteForce(n int) int {
	sum := 0
	sumOfSquares := 0
	for i := 1; i <= n; i++ {
		sum += i
		sumOfSquares += i * i
	}
	return sum*sum - sumOfSquares
}

// gaussFormulas uses closed-form expressions:
// Sum of 1..n = n(n+1)/2
// Sum of squares = n(n+1)(2n+1)/6
func gaussFormulas(n int) int {
	sum := n * (n + 1) / 2
	sumOfSquares := n * (n + 1) * (2*n + 1) / 6
	return sum*sum - sumOfSquares
}

// directFormula uses the algebraic simplification:
// Difference = n(n-1)(n+1)(3n+2)/12
func directFormula(n int) int {
	return n * (n - 1) * (n + 1) * (3*n + 2) / 12
}

// functional uses a reduce-style approach
func functional(n int) int {
	sum := 0
	sumOfSquares := 0
	for i := 1; i <= n; i++ {
		sum += i
		sumOfSquares += i * i
	}
	return sum*sum - sumOfSquares
}

func main() {
	fmt.Println("Problem 006: Sum Square Difference")
	fmt.Printf("Limit: %d\n\n", limit)

	start := time.Now()
	result1 := bruteForce(limit)
	elapsed1 := time.Since(start)

	start = time.Now()
	result2 := gaussFormulas(limit)
	elapsed2 := time.Since(start)

	start = time.Now()
	result3 := directFormula(limit)
	elapsed3 := time.Since(start)

	start = time.Now()
	result4 := functional(limit)
	elapsed4 := time.Since(start)

	fmt.Printf("Brute Force:     %d  (%v)\n", result1, elapsed1)
	fmt.Printf("Gauss Formulas:  %d  (%v)\n", result2, elapsed2)
	fmt.Printf("Direct Formula:  %d  (%v)\n", result3, elapsed3)
	fmt.Printf("Functional:      %d  (%v)\n", result4, elapsed4)

	if result1 != result2 || result2 != result3 || result3 != result4 {
		fmt.Println("\nWARNING: Results do not match!")
	}
}
