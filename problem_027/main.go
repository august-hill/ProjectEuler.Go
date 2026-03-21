// Answer: -59231
// Problem 27: Quadratic Primes
// Find the product of coefficients a and b, for the quadratic n^2 + an + b
// that produces the maximum number of primes for consecutive values of n.

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const primeLimit = 1000000

func sieve(limit int) []bool {
	isPrime := make([]bool, limit)
	for i := 2; i < limit; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i < limit; i++ {
		if isPrime[i] {
			for j := i * i; j < limit; j += i {
				isPrime[j] = false
			}
		}
	}
	return isPrime
}

func solve() int64 {
	isPrime := sieve(primeLimit)

	checkPrime := func(n int) bool {
		if n < 2 || n >= primeLimit {
			return false
		}
		return isPrime[n]
	}

	maxN := 0
	result := 0

	for a := -999; a < 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			n := 0
			for {
				val := n*n + a*n + b
				if val < 0 || !checkPrime(val) {
					break
				}
				n++
			}

			if n > maxN {
				maxN = n
				result = a * b
			}
		}
	}

	return int64(result)
}

func main() { bench.Run(27, solve) }
