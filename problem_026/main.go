// Answer: 983
// Problem 26: Reciprocal Cycles
// Find the value of d < 1000 for which 1/d contains the longest recurring cycle.

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func cycleLength(d int) int {
	seen := make([]int, d)
	for i := range seen {
		seen[i] = -1
	}

	remainder := 1
	position := 0

	for remainder != 0 {
		if seen[remainder] >= 0 {
			return position - seen[remainder]
		}
		seen[remainder] = position
		remainder = (remainder * 10) % d
		position++
	}

	return 0 // Terminating decimal
}

func solve() int64 {
	maxCycle := 0
	result := 0

	for d := 2; d < 1000; d++ {
		cycle := cycleLength(d)
		if cycle > maxCycle {
			maxCycle = cycle
			result = d
		}
	}

	return int64(result)
}

func main() { bench.Run(26, solve) }
