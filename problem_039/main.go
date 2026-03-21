// Problem 039: Integer Right Triangles
// For which value of p <= 1000, is the number of solutions maximised?
// (p = a + b + c for right triangle with integer sides)
// Answer: 840

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	solutions := [1001]int{}

	for a := 1; a < 334; a++ {
		for b := a; b < 500; b++ {
			cSquared := a*a + b*b
			c := int(math.Sqrt(float64(cSquared)))

			if c*c == cSquared {
				p := a + b + c
				if p <= 1000 {
					solutions[p]++
				}
			}
		}
	}

	maxSolutions := 0
	result := 0

	for p := 1; p <= 1000; p++ {
		if solutions[p] > maxSolutions {
			maxSolutions = solutions[p]
			result = p
		}
	}

	return int64(result)
}

func main() { bench.Run(39, solve) }
