// Problem 132: Large Repunit Factors
// Find the sum of the first 40 prime factors of R(10^9).
// Answer: 843296

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func modPow132(base, exp, mod int64) int64 {
	result := int64(1)
	base %= mod
	for exp > 0 {
		if exp&1 != 0 {
			result = result * base % mod
		}
		base = base * base % mod
		exp >>= 1
	}
	return result
}

func isPrime132(n int) bool {
	if n < 2 {
		return false
	}
	if n < 4 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func solve() int64 {
	var sum int64
	count := 0
	for p := 2; count < 40; p++ {
		if !isPrime132(p) {
			continue
		}
		if p == 3 {
			continue
		}
		if modPow132(10, 1000000000, int64(p)) == 1 {
			sum += int64(p)
			count++
		}
	}
	return sum
}

func main() { bench.Run(132, solve) }
