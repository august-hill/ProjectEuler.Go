// Problem 114: Counting Block Combinations I
// Count ways to fill a row of 50 using red blocks of min length 3.
// Answer: 16475640049

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func solve() int64 {
	n := 50
	dp := make([]int64, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]
		for length := 3; length <= i; length++ {
			start := i - length
			if start == 0 {
				dp[i]++
			} else if start == 1 {
				dp[i]++
			} else {
				dp[i] += dp[start-1]
			}
		}
	}
	return dp[n]
}

func main() { bench.Run(114, solve) }
