// Problem 41: Pandigital Prime
// Find the largest n-digit pandigital prime.
// Answer: 7652413

package main

import (
	"fmt"
	"time"
)

// Sieve of Eratosthenes
func sieve(max int) []bool {
	isPrime := make([]bool, max+1)
	for i := 2; i <= max; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= max; i++ {
		if isPrime[i] {
			for j := i * i; j <= max; j += i {
				isPrime[j] = false
			}
		}
	}
	return isPrime
}

// Check if n is pandigital (uses digits 1 to k exactly once where k = digit count)
func isPandigital(n int) bool {
	var digits [10]bool
	digits[0] = true // 0 is not allowed
	digits[8] = true // 8 and 9 pre-marked (we won't see them in valid pandigitals)
	digits[9] = true

	k := 0
	for n > 0 {
		d := n % 10
		if digits[d] {
			return false
		}
		digits[d] = true
		n /= 10
		k++
	}

	for i := 1; i < k; i++ {
		if !digits[i] {
			return false
		}
	}
	return true
}

func solve() int {
	// Max 7-digit pandigital (8,9 digit pandigitals have digit sums divisible by 3)
	const maxPandigital = 7654321
	isPrime := sieve(maxPandigital)

	for n := maxPandigital; n >= 2; n-- {
		if isPrime[n] && isPandigital(n) {
			return n
		}
	}
	return 0
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

	fmt.Println("Problem 41: Pandigital Prime")
	fmt.Println("=============================")
	fmt.Printf("Finding largest pandigital prime, Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
