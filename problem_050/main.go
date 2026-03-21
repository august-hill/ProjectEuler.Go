// Problem 50: Consecutive Prime Sum
// Find prime < 1,000,000 that is sum of most consecutive primes.
// Answer: 997651

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

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

func solve() int64 {
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
	return int64(maxSum)
}

func main() { bench.Run(50, solve) }
