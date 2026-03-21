// Problem 094: Almost Equilateral Triangles
// Find the sum of perimeters of all almost equilateral triangles
// (sides a,a,a+/-1) with integral area and perimeters <= 1,000,000,000.
// Answer: 518408346

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	var limit int64 = 1000000000
	var total int64

	// Case 1: sides a, a, a+1. Perimeter = 3a+1.
	// Recurrence: a_{n+1} = 14*a_n - a_{n-1} - 4
	{
		var aPrev, aCurr int64 = 1, 5
		for {
			perimeter := 3*aCurr + 1
			if perimeter > limit {
				break
			}
			total += perimeter
			aNext := 14*aCurr - aPrev - 4
			aPrev = aCurr
			aCurr = aNext
		}
	}

	// Case 2: sides a, a, a-1. Perimeter = 3a-1.
	// Recurrence: a_{n+1} = 14*a_n - a_{n-1} + 4
	{
		var aPrev, aCurr int64 = 1, 17
		for {
			perimeter := 3*aCurr - 1
			if perimeter > limit {
				break
			}
			total += perimeter
			aNext := 14*aCurr - aPrev + 4
			aPrev = aCurr
			aCurr = aNext
		}
	}

	return total
}

func main() { bench.Run(94, solve) }
