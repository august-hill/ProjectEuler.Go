// Problem 100: Arranged Probability
// Find the number of blue discs for the first arrangement with over 10^12 discs
// where P(two blue) = 1/2.
// Answer: 756872327473

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	// Recurrence: b_{k+1} = 3b + 2n - 2, n_{k+1} = 4b + 3n - 3
	var limit int64 = 1000000000000
	var b, n int64 = 15, 21

	for n <= limit {
		newB := 3*b + 2*n - 2
		newN := 4*b + 3*n - 3
		b = newB
		n = newN
	}
	return b
}

func main() { bench.Run(100, solve) }
