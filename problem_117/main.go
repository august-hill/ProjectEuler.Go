// Problem 117: Red, Green, and Blue Tiles
// Count ways to fill a row of 50 using grey, red (2), green (3), and blue (4) tiles.
// Answer: 100808458960497

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func solve() int64 {
	n := 50
	dp := make([]int64, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]
		if i >= 2 {
			dp[i] += dp[i-2]
		}
		if i >= 3 {
			dp[i] += dp[i-3]
		}
		if i >= 4 {
			dp[i] += dp[i-4]
		}
	}
	return dp[n]
}

func main() { bench.Run(117, solve) }
