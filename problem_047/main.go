// Problem 47: Distinct Prime Factors
// Find first of 4 consecutive integers each with 4 distinct prime factors.
// Answer: 134043

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func countDistinctPrimeFactors(n int, primes []int) int {
	count := 0
	for _, p := range primes {
		if p*p > n {
			break
		}
		if n%p == 0 {
			count++
			for n%p == 0 {
				n /= p
			}
		}
	}
	if n > 1 {
		count++ // remaining factor is prime
	}
	return count
}

func sievePrimes(max int) []int {
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
	return primes
}

func solve() int64 {
	const limit = 150000
	const consecutive = 4
	primes := sievePrimes(limit / 2)

	count := 0
	for i := 2; i < limit; i++ {
		if countDistinctPrimeFactors(i, primes) == consecutive {
			count++
			if count == consecutive {
				return int64(i - consecutive + 1)
			}
		} else {
			count = 0
		}
	}
	return 0
}

func main() { bench.Run(47, solve) }
