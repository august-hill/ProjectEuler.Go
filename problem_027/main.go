// Problem 27: Quadratic Primes
// Find the product of coefficients a and b, for the quadratic n^2 + an + b
// that produces the maximum number of primes for consecutive values of n.
// Answer: -59231

package main

import (
	"fmt"
	"time"
)

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

func solve() int {
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

	return result
}

func benchmark(iterations int) time.Duration {
	// Warmup
	for i := 0; i < 10; i++ {
		solve()
	}

	start := time.Now()
	var result int
	for i := 0; i < iterations; i++ {
		result = solve()
	}
	elapsed := time.Since(start)
	fmt.Printf("Result: %d (%.2f ns/op)\n", result, float64(elapsed.Nanoseconds())/float64(iterations))
	return elapsed
}

func main() {
	const iterations = 10

	fmt.Println("Problem 27: Quadratic Primes")
	fmt.Println("=============================")
	fmt.Printf("Finding product a*b for n^2 + an + b, Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
