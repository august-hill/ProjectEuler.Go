// Problem 058: Spiral Primes
// What is the side length of the square spiral for which the ratio of primes
// along both diagonals first falls below 10%?
// Answer: 26241

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func solve() int64 {
	primeCount := 0
	totalDiagonals := 1 // center

	for n := 1; ; n++ {
		side := 2*n + 1
		corner := side * side

		// Check 3 corners (skip bottom-right which is always a perfect square)
		for i := 1; i <= 3; i++ {
			val := corner - 2*n*i
			if isPrime(val) {
				primeCount++
			}
		}

		totalDiagonals += 4

		// Check ratio: primes/total < 0.1
		if primeCount*10 < totalDiagonals {
			return int64(side)
		}
	}
}

func main() { bench.Run(58, solve) }
