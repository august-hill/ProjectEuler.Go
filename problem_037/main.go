package main

import (
	"fmt"
	"time"
)

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

func solve() int {
	isPrime := sieve(LIMIT)
	sum := 0
	count := 0

	for p := 11; p < LIMIT && count < 11; p++ {
		if isPrime[p] && isLeftTruncatablePrime(p, isPrime) && isRightTruncatablePrime(p, isPrime) {
			sum += p
			count++
		}
	}
	return sum
}

func main() {
	const iterations = 100

	fmt.Println("Problem 37: Truncatable Primes")
	fmt.Println("===============================")

	// Warmup
	for i := 0; i < 10; i++ {
		solve()
	}

	// Benchmark
	start := time.Now()
	var result int
	for i := 0; i < iterations; i++ {
		result = solve()
	}
	elapsed := time.Since(start)

	fmt.Printf("Result: %d (%.2f µs/op)\n", result, float64(elapsed.Microseconds())/float64(iterations))
}
