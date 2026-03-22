// Problem 175: Fractions involving the number of different ways a number can be expressed
// as a sum of powers of 2
// Answer: 13717429 (sum of run lengths in Shortened Binary Expansion "1,13717420,8")

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	p := int64(123456789)
	q := int64(987654321)

	// Reduce by gcd
	a, b := p, q
	for b != 0 {
		a, b = b, a%b
	}
	p /= a
	q /= a

	// Generate runs (LSB first)
	runs := make([]int64, 0, 100)

	for p > 0 && q > 0 {
		if p <= q {
			var k int64
			if p == q {
				k = 1
				q = 0
			} else {
				k = (q - 1) / p
				q -= k * p
				if q == p {
					k++
					q = 0
				}
			}
			runs = append(runs, k)
		} else {
			k := (p - 1) / q
			p -= k * q
			runs = append(runs, k)
		}
	}

	// Sum of all run lengths
	sumRuns := int64(0)
	for _, r := range runs {
		sumRuns += r
	}
	return sumRuns
}

func main() { bench.Run(175, solve) }
