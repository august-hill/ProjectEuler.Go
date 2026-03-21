// Problem 49: Prime Permutations
// Find 4-digit arithmetic sequence of 3 primes that are permutations (not 1487).
// Answer: 296962999629

package main

import (
	"sort"
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

func sortedDigits(n int) string {
	s := strconv.Itoa(n)
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}

func solve() int64 {
	isPrime := sieve(10000)

	for p := 1000; p < 10000; p++ {
		if !isPrime[p] || p == 1487 {
			continue
		}

		sig := sortedDigits(p)
		var perms []int

		// Find all 4-digit prime permutations
		for q := p; q < 10000; q++ {
			if isPrime[q] && sortedDigits(q) == sig {
				perms = append(perms, q)
			}
		}

		// Check for arithmetic sequences
		for i := 0; i < len(perms); i++ {
			for j := i + 1; j < len(perms); j++ {
				diff := perms[j] - perms[i]
				third := perms[j] + diff
				for k := j + 1; k < len(perms); k++ {
					if perms[k] == third {
						// Found it! Return concatenation
						return int64(perms[i])*100000000 + int64(perms[j])*10000 + int64(third)
					}
				}
			}
		}
	}
	return 0
}

func main() { bench.Run(49, solve) }
