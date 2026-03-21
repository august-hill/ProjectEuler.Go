// Answer: 104743
// Problem 007: 10001st Prime
// Find the 10,001st prime number.

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const target = 10001

// sieve generates primes up to limit using Sieve of Eratosthenes
func sieve(limit int) []int {
	isComposite := make([]bool, limit+1)
	for i := 2; i*i <= limit; i++ {
		if !isComposite[i] {
			for j := i * i; j <= limit; j += i {
				isComposite[j] = true
			}
		}
	}
	primes := []int{}
	for i := 2; i <= limit; i++ {
		if !isComposite[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

// sieveNth uses sieve with estimated upper bound
func solve() int64 {
	n := target
	limit := n * 15 // Safe for n up to ~100,000
	if n < 6 {
		limit = 15
	}
	primes := sieve(limit)
	return int64(primes[n-1])
}

func main() { bench.Run(7, solve) }
