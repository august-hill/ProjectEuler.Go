// Problem 191: Prize Strings
// Answer: 1918080160

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	// dp[day][consecutive_absences][late_count]
	var dp [31][3][2]int64
	dp[0][0][0] = 1

	for day := 0; day < 30; day++ {
		for a := 0; a < 3; a++ {
			for l := 0; l < 2; l++ {
				if dp[day][a][l] == 0 {
					continue
				}
				val := dp[day][a][l]
				// On time
				dp[day+1][0][l] += val
				// Late
				if l < 1 {
					dp[day+1][0][l+1] += val
				}
				// Absent
				if a < 2 {
					dp[day+1][a+1][l] += val
				}
			}
		}
	}

	total := int64(0)
	for a := 0; a < 3; a++ {
		for l := 0; l < 2; l++ {
			total += dp[30][a][l]
		}
	}
	return total
}

func main() { bench.Run(191, solve) }
