// Problem 76: Counting Summations
// How many different ways can 100 be written as a sum of at least two positive integers?
// Answer: 190569291

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	target := 100
	dp := make([]int, target+1)
	dp[0] = 1

	for part := 1; part < target; part++ {
		for i := part; i <= target; i++ {
			dp[i] += dp[i-part]
		}
	}

	return int64(dp[target])
}

func main() { bench.Run(76, solve) }
