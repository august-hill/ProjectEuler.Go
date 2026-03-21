// Problem 81: Path Sum: Two Ways
// Find the minimal path sum from top-left to bottom-right, moving only right and down.
// Answer: 427337

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

var (
	matrix081  [80][80]int
	rows081    int
	cols081    int
	loadOnce081 sync.Once
)

func loadData081() {
	loadOnce081.Do(func() {
		data, err := os.ReadFile("p081_matrix.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not open p081_matrix.txt\n")
			os.Exit(1)
		}

		lines := strings.Split(strings.TrimSpace(string(data)), "\n")
		rows081 = 0
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if len(line) < 2 {
				continue
			}
			cols081 = 0
			for _, tok := range strings.Split(line, ",") {
				v, _ := strconv.Atoi(strings.TrimSpace(tok))
				matrix081[rows081][cols081] = v
				cols081++
			}
			rows081++
		}
	})
}

func solve() int64 {
	loadData081()

	dp := [80][80]int{}
	dp[0][0] = matrix081[0][0]
	for j := 1; j < cols081; j++ {
		dp[0][j] = dp[0][j-1] + matrix081[0][j]
	}
	for i := 1; i < rows081; i++ {
		dp[i][0] = dp[i-1][0] + matrix081[i][0]
	}
	for i := 1; i < rows081; i++ {
		for j := 1; j < cols081; j++ {
			up := dp[i-1][j]
			left := dp[i][j-1]
			if up < left {
				dp[i][j] = up + matrix081[i][j]
			} else {
				dp[i][j] = left + matrix081[i][j]
			}
		}
	}
	return int64(dp[rows081-1][cols081-1])
}

func main() { bench.Run(81, solve) }
