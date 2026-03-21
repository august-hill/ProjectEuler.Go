// Problem 37: Truncatable Primes
// Find the sum of the only eleven primes that are both truncatable from left to right and right to left.
// Answer: 748317

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const LIMIT = 1000000

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

func isLeftTruncatablePrime(n int, isPrime []bool) bool {
	// Remove digits from left: 3797 -> 797 -> 97 -> 7
	divisor := 1
	for divisor <= n {
		divisor *= 10
	}
	divisor /= 10

	for divisor > 1 {
		n = n % divisor
		if n == 0 || !isPrime[n] {
			return false
		}
		divisor /= 10
	}
	return true
}

func isRightTruncatablePrime(n int, isPrime []bool) bool {
	// Remove digits from right: 3797 -> 379 -> 37 -> 3
	n /= 10
	for n > 0 {
		if !isPrime[n] {
			return false
		}
		n /= 10
	}
	return true
}

func solve() int64 {
	isPrime := sieve(LIMIT)
	sum := 0
	count := 0

	for p := 11; p < LIMIT && count < 11; p++ {
		if isPrime[p] && isLeftTruncatablePrime(p, isPrime) && isRightTruncatablePrime(p, isPrime) {
			sum += p
			count++
		}
	}
	return int64(sum)
}

func main() { bench.Run(37, solve) }
