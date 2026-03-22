// Problem 133: Repunit Nonfactors
// Find the sum of primes below 100000 that never divide R(10^n).
// Answer: 453647705

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const limit133 = 100000

func modPow133(base, exp, mod int64) int64 {
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

func solve() int64 {
	notPrime := make([]bool, limit133)
	notPrime[0] = true
	notPrime[1] = true
	for i := 2; i*i < limit133; i++ {
		if !notPrime[i] {
			for j := i * i; j < limit133; j += i {
				notPrime[j] = true
			}
		}
	}

	// 10^16 = 2^16 * 5^16, covers all A(p) of the form 2^a * 5^b up to p < 100000
	var exp int64 = 1
	for i := 0; i < 16; i++ {
		exp *= 10
	}

	var sum int64
	for p := 2; p < limit133; p++ {
		if notPrime[p] {
			continue
		}
		if p == 2 || p == 5 {
			sum += int64(p)
			continue
		}
		if p == 3 {
			sum += 3
			continue
		}
		if modPow133(10, exp, int64(p)) != 1 {
			sum += int64(p)
		}
	}
	return sum
}

func main() { bench.Run(133, solve) }
