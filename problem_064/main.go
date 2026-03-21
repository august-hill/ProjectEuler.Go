// Problem 064: Odd Period Square Roots
// How many continued fractions for sqrt(N), N <= 10000, have an odd period?
// Answer: 1322

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	count := 0

	for n := 2; n <= 10000; n++ {
		a0 := int(math.Sqrt(float64(n)))
		if a0*a0 == n {
			continue
		}

		m, d, a := 0, 1, a0
		period := 0

		for {
			m = d*a - m
			d = (n - m*m) / d
			a = (a0 + m) / d
			period++
			if a == 2*a0 {
				break
			}
		}

		if period%2 == 1 {
			count++
		}
	}

	return int64(count)
}

func main() { bench.Run(64, solve) }
