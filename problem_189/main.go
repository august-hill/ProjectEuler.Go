// Problem 189: Tri-colouring a Triangular Grid
// Answer: 10834893628237824

package main

import (
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const maxRow189 = 8

var (
	once189        sync.Once
	answerCache189 int64
	pow3_189       [10]int
)

func initPow3189() {
	pow3_189[0] = 1
	for i := 1; i <= 9; i++ {
		pow3_189[i] = pow3_189[i-1] * 3
	}
}

func getColour189(state, pos int) int {
	return (state / pow3_189[pos]) % 3
}

func countDownWays189(prevUp []int, r int, curUp []int) int64 {
	ways := int64(1)
	for j := 0; j < r; j++ {
		a, b, c := prevUp[j], curUp[j], curUp[j+1]
		distinct := 1
		if b != a {
			distinct++
		}
		if c != a && c != b {
			distinct++
		}
		if distinct == 3 {
			return 0
		}
		ways *= int64(3 - distinct)
	}
	return ways
}

func compute189() {
	initPow3189()

	dp := make([]int64, 6561) // 3^8
	ndp := make([]int64, 6561)

	// Row 1: 1 up triangle, 3 choices
	dp[0] = 1
	dp[1] = 1
	dp[2] = 1

	prevUp := make([]int, maxRow189+1)
	curUp := make([]int, maxRow189+1)

	for row := 1; row < maxRow189; row++ {
		nupCur := row
		nupNext := row + 1

		for i := range ndp {
			ndp[i] = 0
		}

		for s := 0; s < pow3_189[nupCur]; s++ {
			if dp[s] == 0 {
				continue
			}
			for j := 0; j < nupCur; j++ {
				prevUp[j] = getColour189(s, j)
			}

			for ns := 0; ns < pow3_189[nupNext]; ns++ {
				for j := 0; j < nupNext; j++ {
					curUp[j] = getColour189(ns, j)
				}
				ways := countDownWays189(prevUp, nupCur, curUp)
				if ways > 0 {
					ndp[ns] += dp[s] * ways
				}
			}
		}
		dp, ndp = ndp, dp
	}

	total := int64(0)
	for s := 0; s < pow3_189[maxRow189]; s++ {
		total += dp[s]
	}
	answerCache189 = total
}

func solve() int64 {
	once189.Do(compute189)
	return answerCache189
}

func main() { bench.Run(189, solve) }
