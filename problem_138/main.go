// Problem 138: Special Isosceles Triangles
// Find the sum of the first 12 L values for special isosceles triangles.
// Answer: 1118049290473932

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func solve() int64 {
	var sum int64
	lPrev := int64(1)
	lCurr := int64(17)

	for i := 0; i < 12; i++ {
		sum += lCurr
		lNext := 18*lCurr - lPrev
		lPrev = lCurr
		lCurr = lNext
	}
	return sum
}

func main() { bench.Run(138, solve) }
