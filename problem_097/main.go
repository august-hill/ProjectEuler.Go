// Problem 097: Large Non-Mersenne Prime
// Find the last ten digits of 28433 * 2^7830457 + 1.
// Answer: 8739992577

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

const mod097 = 10000000000 // 10^10

func modPow097(base, exp, m uint64) uint64 {
	result := uint64(1)
	base = base % m
	for exp > 0 {
		if exp&1 == 1 {
			result = mulmod097(result, base, m)
		}
		base = mulmod097(base, base, m)
		exp >>= 1
	}
	return result
}

func mulmod097(a, b, m uint64) uint64 {
	var result uint64
	a = a % m
	b = b % m
	for b > 0 {
		if b&1 == 1 {
			result = (result + a) % m
		}
		a = (a * 2) % m
		b >>= 1
	}
	return result
}

func solve() int64 {
	power := modPow097(2, 7830457, mod097)
	result := mulmod097(power, 28433, mod097)
	result = (result + 1) % mod097
	return int64(result)
}

func main() { bench.Run(97, solve) }
