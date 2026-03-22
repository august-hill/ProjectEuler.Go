// Problem 116: Red, Green or Blue Tiles
// Count ways to replace grey tiles in a row of 50 with colored tiles (no mixing).
// Answer: 20492570929

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func countWays116(tileLen, n int) int64 {
	dp := make([]int64, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]
		if i >= tileLen {
			dp[i] += dp[i-tileLen]
		}
	}
	return dp[n] - 1 // subtract all-grey
}

func solve() int64 {
	n := 50
	return countWays116(2, n) + countWays116(3, n) + countWays116(4, n)
}

func main() { bench.Run(116, solve) }
