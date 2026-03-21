// Problem 85: Counting Rectangles
// Find the area of the grid with the nearest solution to containing two million rectangles.
// Answer: 2772

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	target := int64(2000000)
	bestArea := 0
	bestDiff := target

	for m := 1; m <= 2000; m++ {
		cm := int64(m) * int64(m+1) / 2
		if cm > target {
			break
		}
		for n := m; n <= 2000; n++ {
			count := cm * int64(n) * int64(n+1) / 2
			diff := count - target
			if diff < 0 {
				diff = -diff
			}
			if diff < bestDiff {
				bestDiff = diff
				bestArea = m * n
			}
			if count > target {
				break
			}
		}
	}
	return int64(bestArea)
}

func main() { bench.Run(85, solve) }
