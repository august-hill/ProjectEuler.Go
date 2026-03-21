// Problem 73: Counting Fractions in a Range
// How many fractions lie between 1/3 and 1/2 in the sorted set of reduced proper fractions for d <= 12,000?
// Answer: 7295372

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func gcd073(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func solve() int64 {
	limit := 12000
	count := 0

	for d := 2; d <= limit; d++ {
		nMin := d/3 + 1
		var nMax int
		if d%2 == 0 {
			nMax = d/2 - 1
		} else {
			nMax = d / 2
		}

		for n := nMin; n <= nMax; n++ {
			if gcd073(n, d) == 1 {
				count++
			}
		}
	}

	return int64(count)
}

func main() { bench.Run(73, solve) }
