// Problem 36: Double-base Palindromes
// Find the sum of all numbers less than one million which are palindromic in both base 10 and base 2.
// Answer: 872187

package main

import (
	"fmt"
	"strconv"
	"time"
)

func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func solve() int {
	sum := 0

	for n := 1; n < 1000000; n++ {
		dec := strconv.Itoa(n)
		bin := strconv.FormatInt(int64(n), 2)

		if isPalindrome(dec) && isPalindrome(bin) {
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

	fmt.Println("Problem 36: Double-base Palindromes")
	fmt.Println("=====================================")
	fmt.Printf("Sum of double-base palindromes below 1 million, Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
