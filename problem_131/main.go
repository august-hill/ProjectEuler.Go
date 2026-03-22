// Problem 131: Prime Cube Partnership
// Count primes p < 10^6 where p = (b^3 - a^3) for consecutive a, b.
// Answer: 173

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func isPrime131(n int64) bool {
	if n < 2 {
		return false
	}
	if n < 4 {
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

func solve() int64 {
	count := 0
	for a := int64(1); ; a++ {
		p := 3*a*a + 3*a + 1
		if p >= 1000000 {
			break
		}
		if isPrime131(p) {
			count++
		}
	}
	return int64(count)
}

func main() { bench.Run(131, solve) }
