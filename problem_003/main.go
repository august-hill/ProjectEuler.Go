// Answer: 6857
// Problem 003: Largest Prime Factor
// Find the largest prime factor of 600851475143.

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const target = 600851475143

// optimizedTrialDivision uses 6k+/-1 optimization
func solve() int64 {
	n := int64(target)
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

	// Check factors of form 6k+/-1
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

func main() { bench.Run(3, solve) }
