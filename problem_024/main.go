// Problem 24: Lexicographic Permutations
// What is the millionth lexicographic permutation of the digits 0-9?
// Answer: 2783915460

package main

import (
	"fmt"
	"time"
)

func factorial(n int) int64 {
	result := int64(1)
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}
	return result
}

func solve() int64 {
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := int64(999999) // 0-indexed
	result := int64(0)

	for i := 9; i >= 1; i-- {
		fact := factorial(i)
		idx := int(n / fact)
		result = result*10 + int64(digits[idx])

		// Remove digit at idx
		digits = append(digits[:idx], digits[idx+1:]...)
		n %= fact
	}
	result = result*10 + int64(digits[0])

	return result
}

func benchmark(iterations int) time.Duration {
	// Warmup
	for i := 0; i < 10; i++ {
		solve()
	}

	start := time.Now()
	var result int64
	for i := 0; i < iterations; i++ {
		result = solve()
	}
	elapsed := time.Since(start)
	fmt.Printf("Result: %d (%.2f ns/op)\n", result, float64(elapsed.Nanoseconds())/float64(iterations))
	return elapsed
}

func main() {
	const iterations = 10000

	fmt.Println("Problem 24: Lexicographic Permutations")
	fmt.Println("=======================================")
	fmt.Printf("Finding millionth lexicographic permutation, Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
