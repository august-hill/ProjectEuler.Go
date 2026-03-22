// Problem 164: Numbers for which no three consecutive digits have a sum > 9
// Answer: 378158756814587

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	var dp [10][10]int64

	// Place first two digits
	for d1 := 1; d1 <= 9; d1++ {
		for d2 := 0; d2 <= 9; d2++ {
			dp[d1][d2] = 1
		}
	}

	// Extend for digits 3 through 20
	for pos := 3; pos <= 20; pos++ {
		var ndp [10][10]int64
		for d1 := 0; d1 <= 9; d1++ {
			for d2 := 0; d2 <= 9; d2++ {
				if dp[d1][d2] == 0 {
					continue
				}
				maxD3 := 9 - d1 - d2
				if maxD3 < 0 {
					continue
				}
				for d3 := 0; d3 <= maxD3; d3++ {
					ndp[d2][d3] += dp[d1][d2]
				}
			}
		}
		dp = ndp
	}

	total := int64(0)
	for d1 := 0; d1 <= 9; d1++ {
		for d2 := 0; d2 <= 9; d2++ {
			total += dp[d1][d2]
		}
	}
	return total
}

func main() { bench.Run(164, solve) }
