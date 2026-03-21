// Answer: 142913828922
// Problem 10: Summation of Primes
// Find the sum of all primes below two million.

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const limit = 2_000_000

// sieve returns the sum of all primes below n using Sieve of Eratosthenes
func solve() int64 {
	n := limit
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

func main() { bench.Run(10, solve) }
