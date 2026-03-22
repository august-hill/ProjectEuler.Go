// Problem 135: Same Differences
// Find how many values of n < 10^6 have exactly 10 solutions.
// Answer: 4989

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const limit135 = 1000000

func solve() int64 {
	count := make([]int, limit135)

	for y := 1; y < limit135; y++ {
		dMin := y/4 + 1
		dMax := y - 1
		maxN := int64(limit135 - 1)
		dLim := int((maxN + int64(y)*int64(y)) / (4 * int64(y)))
		if dLim < dMax {
			dMax = dLim
		}

		for d := dMin; d <= dMax; d++ {
			n := y * (4*d - y)
			if n > 0 && n < limit135 {
				count[n]++
			}
		}
	}

	result := 0
	for n := 1; n < limit135; n++ {
		if count[n] == 10 {
			result++
		}
	}
	return int64(result)
}

func main() { bench.Run(135, solve) }
