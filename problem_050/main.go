// Problem 50: Consecutive Prime Sum
// Find prime < 1,000,000 that is sum of most consecutive primes.
// Answer: 997651

package main

import (
	"fmt"
	"time"
)

func sievePrimes(max int) ([]bool, []int) {
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
	var primes []int
	for i := 2; i <= max; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return isPrime, primes
}

func solve() int {
	const limit = 1000000
	isPrime, primes := sievePrimes(limit)

	maxLen := 0
	maxSum := 0

	for i := 0; i < len(primes); i++ {
		sum := 0
		for j := 0; i+j < len(primes); j++ {
			sum += primes[i+j]
			if sum >= limit {
				break
			}
			if isPrime[sum] && j > maxLen {
				maxLen = j
				maxSum = sum
			}
		}
	}
	return maxSum
}

func benchmark(iterations int) time.Duration {
	for i := 0; i < 10; i++ {
		solve()
	}
	start := time.Now()
	var result int
	for i := 0; i < iterations; i++ {
		result = solve()
	}
	elapsed := time.Since(start)
	fmt.Printf("Result: %d (%.2f ms/op)\n", result, float64(elapsed.Milliseconds())/float64(iterations))
	return elapsed
}

func main() {
	fmt.Println("Problem 50: Consecutive Prime Sum")
	fmt.Println("==================================")
	benchmark(100)
}
