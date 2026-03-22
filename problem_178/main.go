// Problem 178: Step Numbers
// Answer: 126461847755

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	var dp [10][1024]int64
	fullMask := (1 << 10) - 1
	total := int64(0)

	// Length 1
	for d := 1; d <= 9; d++ {
		dp[d][1<<d] = 1
	}

	for length := 2; length <= 40; length++ {
		var ndp [10][1024]int64
		for d := 0; d <= 9; d++ {
			for mask := 0; mask < 1024; mask++ {
				if dp[d][mask] == 0 {
					continue
				}
				val := dp[d][mask]
				if d > 0 {
					ndp[d-1][mask|(1<<(d-1))] += val
				}
				if d < 9 {
					ndp[d+1][mask|(1<<(d+1))] += val
				}
			}
		}
		for d := 0; d <= 9; d++ {
			total += ndp[d][fullMask]
		}
		dp = ndp
	}
	return total
}

func main() { bench.Run(178, solve) }
