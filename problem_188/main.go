// Problem 188: The Hyperexponentiation of a Number
// Answer: 95962097

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

const mod188 = 100000000

func modPow188(base, exp, mod uint64) uint64 {
	result := uint64(1)
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			result = result * base % mod
		}
		exp >>= 1
		base = base * base % mod
	}
	return result
}

func eulerTotient188(n uint64) uint64 {
	result := n
	orig := n
	for p := uint64(2); p*p <= orig; p++ {
		if n%p == 0 {
			for n%p == 0 {
				n /= p
			}
			result -= result / p
		}
	}
	if n > 1 {
		result -= result / n
	}
	return result
}

func hyper188(a, b, m uint64) uint64 {
	if m == 1 {
		return 0
	}
	if b == 1 {
		return a % m
	}
	phi := eulerTotient188(m)
	exp := hyper188(a, b-1, phi)
	return modPow188(a, exp+phi, m)
}

func solve() int64 {
	return int64(hyper188(1777, 1855, mod188))
}

func main() { bench.Run(188, solve) }
