// Problem 46: Goldbach's Other Conjecture
// Find smallest odd composite that cannot be written as prime + 2*square.
// Answer: 5777

package main

import (
	"fmt"
	"time"
)

func sieve(max int) []bool {
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
	return isPrime
}

func solve() int {
	const limit = 10000
	isPrime := sieve(limit)

	for c := 9; c < limit; c += 2 {
		if isPrime[c] {
			continue // skip primes, we want composites
		}

		found := false
		for p := 2; p < c && !found; p++ {
			if isPrime[p] {
				for y := 1; !found; y++ {
					z := p + 2*y*y
					if z > c {
						break
					}
					if z == c {
						found = true
					}
				}
			}
		}

		if !found {
			return c
		}
	}
	return 0
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
	fmt.Printf("Result: %d (%.2f ns/op)\n", result, float64(elapsed.Nanoseconds())/float64(iterations))
	return elapsed
}

func main() {
	fmt.Println("Problem 46: Goldbach's Other Conjecture")
	fmt.Println("========================================")
	benchmark(1000)
}
