// Problem 007: 10001st Prime
// Find the 10,001st prime number.

package main

import (
	"fmt"
	"time"
)

const target = 10001

// trialDivision collects primes and tests each odd candidate
func trialDivision(n int) int {
	primes := []int{2}
	candidate := 3
	for len(primes) < n {
		isPrime := true
		for _, p := range primes {
			if p*p > candidate {
				break
			}
			if candidate%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, candidate)
		}
		candidate += 2
	}
	return primes[n-1]
}

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
func sieveNth(n int) int {
	// Upper bound: n * (ln(n) + ln(ln(n))) for n >= 6
	// For n = 10001, this is about 114,320
	// We'll use a safe overestimate
	limit := n * 15 // Safe for n up to ~100,000
	if n < 6 {
		limit = 15
	}
	primes := sieve(limit)
	return primes[n-1]
}

// sixKPlusMinusOne uses the fact that all primes > 3 are 6k±1
func sixKPlusMinusOne(n int) int {
	if n == 1 {
		return 2
	}
	if n == 2 {
		return 3
	}

	primes := []int{2, 3}
	k := 1
	for len(primes) < n {
		// Check 6k - 1
		candidate := 6*k - 1
		if isPrimeWithList(candidate, primes) {
			primes = append(primes, candidate)
			if len(primes) == n {
				break
			}
		}
		// Check 6k + 1
		candidate = 6*k + 1
		if isPrimeWithList(candidate, primes) {
			primes = append(primes, candidate)
		}
		k++
	}
	return primes[n-1]
}

func isPrimeWithList(candidate int, primes []int) bool {
	for _, p := range primes {
		if p*p > candidate {
			break
		}
		if candidate%p == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Problem 007: 10001st Prime")
	fmt.Printf("Target: %dth prime\n\n", target)

	start := time.Now()
	result1 := trialDivision(target)
	elapsed1 := time.Since(start)

	start = time.Now()
	result2 := sieveNth(target)
	elapsed2 := time.Since(start)

	start = time.Now()
	result3 := sixKPlusMinusOne(target)
	elapsed3 := time.Since(start)

	fmt.Printf("Trial Division:  %d  (%v)\n", result1, elapsed1)
	fmt.Printf("Sieve:           %d  (%v)\n", result2, elapsed2)
	fmt.Printf("6k±1:            %d  (%v)\n", result3, elapsed3)

	if result1 != result2 || result2 != result3 {
		fmt.Println("\nWARNING: Results do not match!")
	}
}
