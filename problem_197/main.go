// Problem 197: Investigating the Behaviour of a Recursively Defined Sequence
// Answer: 1710637717 (representing 1.710637717)

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func f197(x float64) float64 {
	return math.Floor(math.Pow(2.0, 30.403243784-x*x)) * 1e-9
}

func solve() int64 {
	u := -1.0
	for i := 0; i < 1000; i++ {
		u = f197(u)
	}
	sum := u + f197(u)
	return int64(math.Round(sum * 1e9))
}

func main() { bench.Run(197, solve) }
