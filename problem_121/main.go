// Problem 121: Disc Game Prize Fund
// Find the maximum prize fund for the 15-turn disc game.
// Answer: 2269

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func solve() int64 {
	n := 15
	dp := make([]int64, n+1)
	dp[0] = 1

	for k := 1; k <= n; k++ {
		for j := k; j >= 1; j-- {
			dp[j] = dp[j]*int64(k) + dp[j-1]
		}
		dp[0] *= int64(k)
	}

	var denom int64 = 1
	for i := 1; i <= n+1; i++ {
		denom *= int64(i)
	}

	var winNum int64
	for j := n/2 + 1; j <= n; j++ {
		winNum += dp[j]
	}

	return denom / winNum
}

func main() { bench.Run(121, solve) }
