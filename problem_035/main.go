// Problem 35: Circular Primes
// How many circular primes are there below one million?
// Answer: 55

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const circLimit = 1000000

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

func numDigits(n int) int {
	count := 0
	for n > 0 {
		count++
		n /= 10
	}
	return count
}

func rotate(n int) int {
	digits := numDigits(n)
	last := n % 10
	rest := n / 10

	multiplier := 1
	for i := 1; i < digits; i++ {
		multiplier *= 10
	}

	return last*multiplier + rest
}

func isCircularPrime(n int, isPrime []bool) bool {
	digits := numDigits(n)
	rotated := n

	for i := 0; i < digits; i++ {
		if rotated >= circLimit || !isPrime[rotated] {
			return false
		}
		rotated = rotate(rotated)
	}
	return true
}

func solve() int64 {
	isPrime := sieve(circLimit)
	count := 0

	for n := 2; n < circLimit; n++ {
		if isPrime[n] && isCircularPrime(n, isPrime) {
			count++
		}
	}

	return int64(count)
}

func main() { bench.Run(35, solve) }
