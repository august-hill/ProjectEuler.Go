// Problem 182: RSA Encryption
// Answer: 399788195976

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func gcd182(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func solve() int64 {
	p := int64(1009)
	q := int64(3643)
	phi := (p - 1) * (q - 1)

	sum := int64(0)
	for e := int64(2); e < phi; e++ {
		if gcd182(e, phi) != 1 {
			continue
		}
		if gcd182(e-1, p-1) == 2 && gcd182(e-1, q-1) == 2 {
			sum += e
		}
	}
	return sum
}

func main() { bench.Run(182, solve) }
