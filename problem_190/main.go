// Problem 190: Maximising a Weighted Product
// Answer: 371048281

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	total := int64(0)
	for m := 2; m <= 15; m++ {
		logPm := 0.0
		for k := 1; k <= m; k++ {
			logPm += float64(k) * math.Log(2.0*float64(k)/float64(m+1))
		}
		total += int64(math.Floor(math.Exp(logPm)))
	}
	return total
}

func main() { bench.Run(190, solve) }
