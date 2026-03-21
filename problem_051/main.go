// Problem 51: Prime Digit Replacements
// Find smallest prime where replacing 3 digits gives 8 primes.
// Answer: 121313

package main

import (
	"strconv"

	"github.com/august-hill/ProjectEuler.Go/bench"
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

func solve() int64 {
	const upper = 999999
	const lower = upper / 10
	isPrime := sieve(upper)

	for p := lower + 1; p <= upper; p++ {
		if !isPrime[p] {
			continue
		}

		s := strconv.Itoa(p)
		n := len(s)

		// Try all combinations of 3 positions
		for i := 0; i < n-2; i++ {
			for j := i + 1; j < n-1; j++ {
				for k := j + 1; k < n; k++ {
					// Check if these positions have the same digit
					if s[i] != s[j] || s[j] != s[k] {
						continue
					}

					// Try replacing with 0-9
					var primeCount int
					var firstPrime int
					for d := byte('0'); d <= '9'; d++ {
						bs := []byte(s)
						bs[i], bs[j], bs[k] = d, d, d
						num, _ := strconv.Atoi(string(bs))
						if num > lower && isPrime[num] {
							primeCount++
							if firstPrime == 0 {
								firstPrime = num
							}
						}
					}

					if primeCount == 8 {
						return int64(firstPrime)
					}
				}
			}
		}
	}
	return 0
}

func main() { bench.Run(51, solve) }
