// Problem 136: Singleton Difference
// Find how many values of n < 50,000,000 have exactly one solution.
// Answer: 2544559

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const limit136 = 50000000

func solve() int64 {
	count := make([]uint8, limit136)

	for y := 1; y < limit136; y++ {
		dMin := y/4 + 1
		dMax := y - 1
		maxD := (int64(limit136-1) + int64(y)*int64(y)) / (4 * int64(y))
		if int(maxD) < dMax {
			dMax = int(maxD)
		}

		for d := dMin; d <= dMax; d++ {
			n := int64(y) * int64(4*d-y)
			if n > 0 && n < limit136 {
				if count[n] < 3 {
					count[n]++
				}
			}
		}
	}

	result := 0
	for n := 1; n < limit136; n++ {
		if count[n] == 1 {
			result++
		}
	}
	return int64(result)
}

func main() { bench.Run(136, solve) }
