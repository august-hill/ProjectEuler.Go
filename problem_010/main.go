// Problem 10: Summation of Primes
// Find the sum of all primes below two million.

package main

import (
	"fmt"
	"time"
)

const limit = 2_000_000

// sieve returns the sum of all primes below n using Sieve of Eratosthenes
func sieve(n int) int64 {
	if n < 2 {
		return 0
	}

	// Create boolean slice, true = prime candidate
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}

	// Mark composites
	for i := 2; i*i < n; i++ {
		if isPrime[i] {
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	// Sum primes
	var sum int64
	for i := 2; i < n; i++ {
		if isPrime[i] {
			sum += int64(i)
		}
	}
	return sum
}

func benchmark(name string, f func(int) int64, n int, iterations int) time.Duration {
	// Warmup
	for i := 0; i < 10; i++ {
		f(n)
	}

	start := time.Now()
	for i := 0; i < iterations; i++ {
		f(n)
	}
	elapsed := time.Since(start)
	result := f(n)
	fmt.Printf("%s: %d (%.2f ns/op)\n", name, result, float64(elapsed.Nanoseconds())/float64(iterations))
	return elapsed
}

func main() {
	const iterations = 1000

	fmt.Println("Problem 10: Summation of Primes")
	fmt.Println("===============================")
	fmt.Printf("Limit: %d, Iterations: %d\n\n", limit, iterations)

	benchmark("Sieve", sieve, limit, iterations)
}
