// Problem 172: Investigating numbers with few repeated digits
// Answer: 227485267000992000

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

var C172 [20][20]int64

func initC172() {
	for n := 0; n < 20; n++ {
		C172[n][0] = 1
		for k := 1; k <= n; k++ {
			C172[n][k] = C172[n-1][k-1] + C172[n-1][k]
		}
	}
}

func solve() int64 {
	initC172()

	dp := make([]int64, 20)
	dp[18] = 1

	for d := 0; d < 10; d++ {
		ndp := make([]int64, 20)
		for r := 0; r <= 18; r++ {
			if dp[r] == 0 {
				continue
			}
			for c := 0; c <= 3 && c <= r; c++ {
				ndp[r-c] += dp[r] * C172[r][c]
			}
		}
		dp = ndp
	}
	total := dp[0]

	// Subtract sequences starting with 0
	dp2 := make([]int64, 20)
	dp2[17] = 1

	// Digit 0: max 2 more occurrences
	ndp2 := make([]int64, 20)
	for r := 0; r <= 17; r++ {
		if dp2[r] == 0 {
			continue
		}
		for c := 0; c <= 2 && c <= r; c++ {
			ndp2[r-c] += dp2[r] * C172[r][c]
		}
	}
	dp2 = ndp2

	// Digits 1-9: max 3 each
	for d := 1; d < 10; d++ {
		ndp3 := make([]int64, 20)
		for r := 0; r <= 17; r++ {
			if dp2[r] == 0 {
				continue
			}
			for c := 0; c <= 3 && c <= r; c++ {
				ndp3[r-c] += dp2[r] * C172[r][c]
			}
		}
		dp2 = ndp3
	}
	withLeadingZero := dp2[0]

	return total - withLeadingZero
}

func main() { bench.Run(172, solve) }
