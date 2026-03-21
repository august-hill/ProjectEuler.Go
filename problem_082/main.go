// Problem 82: Path Sum: Three Ways
// Find the minimal path sum from left column to right column, moving up, down, and right.
// Answer: 260324

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
	matrix082  [80][80]int
	rows082    int
	cols082    int
	loadOnce082 sync.Once
)

func loadData082() {
	loadOnce082.Do(func() {
		data, err := os.ReadFile("p082_matrix.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not open p082_matrix.txt\n")
			os.Exit(1)
		}

		lines := strings.Split(strings.TrimSpace(string(data)), "\n")
		rows082 = 0
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if len(line) < 2 {
				continue
			}
			cols082 = 0
			for _, tok := range strings.Split(line, ",") {
				v, _ := strconv.Atoi(strings.TrimSpace(tok))
				matrix082[rows082][cols082] = v
				cols082++
			}
			rows082++
		}
	})
}

func solve() int64 {
	loadData082()

	dp := make([]int, rows082)
	newDp := make([]int, rows082)

	for i := 0; i < rows082; i++ {
		dp[i] = matrix082[i][0]
	}

	for j := 1; j < cols082; j++ {
		// From left
		for i := 0; i < rows082; i++ {
			newDp[i] = dp[i] + matrix082[i][j]
		}
		// Pass down
		for i := 1; i < rows082; i++ {
			via := newDp[i-1] + matrix082[i][j]
			if via < newDp[i] {
				newDp[i] = via
			}
		}
		// Pass up
		for i := rows082 - 2; i >= 0; i-- {
			via := newDp[i+1] + matrix082[i][j]
			if via < newDp[i] {
				newDp[i] = via
			}
		}
		copy(dp, newDp)
	}

	minVal := dp[0]
	for i := 1; i < rows082; i++ {
		if dp[i] < minVal {
			minVal = dp[i]
		}
	}
	return int64(minVal)
}

func main() { bench.Run(82, solve) }
