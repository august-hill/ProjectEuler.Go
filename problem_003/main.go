// Problem 003: Largest Prime Factor
// Find the largest prime factor of 600851475143.

package main

import (
	"fmt"
	"math"
	"time"
)

const target = 600851475143

// sieve generates primes up to limit using Sieve of Eratosthenes
func sieve(limit int) []int64 {
	if limit < 2 {
		return []int64{}
	}
	isPrime := make([]bool, limit+1)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0], isPrime[1] = false, false

	for i := 2; i*i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}

	primes := []int64{}
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			primes = append(primes, int64(i))
		}
	}
	return primes
}

// sieveMethod uses precomputed primes up to sqrt(n) for trial division
func sieveMethod(n int64) int64 {
	limit := int(math.Sqrt(float64(n))) + 1
	primes := sieve(limit)

	var largestFactor int64 = 1
	for _, p := range primes {
		for n%p == 0 {
			largestFactor = p
			n /= p
		}
		if n == 1 {
			break
		}
	}
	// If n > 1, then n itself is prime and the largest factor
	if n > 1 {
		largestFactor = n
	}
	return largestFactor
}

// optimizedTrialDivision uses 6k±1 optimization
func optimizedTrialDivision(n int64) int64 {
	var largestFactor int64 = 1

	// Check factor of 2
	for n%2 == 0 {
		largestFactor = 2
		n /= 2
	}

	// Check factor of 3
	for n%3 == 0 {
		largestFactor = 3
		n /= 3
	}

	// Check factors of form 6k±1
	for i := int64(5); i*i <= n; i += 6 {
		for n%i == 0 {
			largestFactor = i
			n /= i
		}
		for n%(i+2) == 0 {
			largestFactor = i + 2
			n /= (i + 2)
		}
	}

	// If n > 1, it's a prime factor
	if n > 1 {
		largestFactor = n
	}
	return largestFactor
}

// gcd computes greatest common divisor
func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// pollardRho finds a non-trivial factor of n (n must be composite)
func pollardRho(n int64) int64 {
	if n%2 == 0 {
		return 2
	}

	x := int64(2)
	y := int64(2)
	d := int64(1)

	// f(x) = x^2 + 1 mod n
	f := func(x int64) int64 {
		return (x*x + 1) % n
	}

	for d == 1 {
		x = f(x)
		y = f(f(y))
		diff := x - y
		if diff < 0 {
			diff = -diff
		}
		d = gcd(diff, n)
	}

	return d
}

// largestPrimeFactorPollard uses Pollard's rho to find largest prime factor
func largestPrimeFactorPollard(n int64) int64 {
	var largestFactor int64 = 1

	// Remove factors of 2
	for n%2 == 0 {
		largestFactor = 2
		n /= 2
	}

	// Remove factors of 3
	for n%3 == 0 {
		largestFactor = 3
		n /= 3
	}

	for n > 1 {
		if isProbablePrime(n) {
			if n > largestFactor {
				largestFactor = n
			}
			break
		}

		factor := pollardRho(n)
		for factor == n {
			// Retry with different starting point would be needed for production
			// For this problem, simple retry works
			factor = pollardRho(n)
		}

		// Factor out this divisor completely
		for n%factor == 0 {
			n /= factor
		}

		// Recursively find largest prime factor of this factor
		pf := largestPrimeFactorPollard(factor)
		if pf > largestFactor {
			largestFactor = pf
		}
	}

	return largestFactor
}

// isProbablePrime does simple primality test (sufficient for this problem size)
func isProbablePrime(n int64) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := int64(5); i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Problem 003: Largest Prime Factor")
	fmt.Printf("Target: %d\n\n", target)

	start := time.Now()
	result1 := sieveMethod(target)
	elapsed1 := time.Since(start)

	start = time.Now()
	result2 := optimizedTrialDivision(target)
	elapsed2 := time.Since(start)

	start = time.Now()
	result3 := largestPrimeFactorPollard(target)
	elapsed3 := time.Since(start)

	fmt.Printf("Sieve + Trial Division:     %d  (%v)\n", result1, elapsed1)
	fmt.Printf("Optimized Trial (6k±1):     %d  (%v)\n", result2, elapsed2)
	fmt.Printf("Pollard's Rho:              %d  (%v)\n", result3, elapsed3)

	if result1 != result2 || result2 != result3 {
		fmt.Println("\nWARNING: Results do not match!")
	}
}
