// Problem 067: Maximum Path Sum II
// Find the maximum total from top to bottom of a 100-row triangle.
// Answer: 7273

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

var triangle [100][100]int
var numRows int
var loadOnce sync.Once

func loadTriangle() {
	loadOnce.Do(func() {
		f, err := os.Open("p067_triangle.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot open p067_triangle.txt: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		numRows = 0
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" {
				continue
			}
			parts := strings.Fields(line)
			for j, p := range parts {
				triangle[numRows][j], _ = strconv.Atoi(p)
			}
			numRows++
		}
	})
}

func solve() int64 {
	loadTriangle()
	var dp [100]int
	// Copy last row
	for j := 0; j < numRows; j++ {
		dp[j] = triangle[numRows-1][j]
	}
	// Work bottom-up
	for i := numRows - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			left := dp[j]
			right := dp[j+1]
			if left > right {
				dp[j] = triangle[i][j] + left
			} else {
				dp[j] = triangle[i][j] + right
			}
		}
	}
	return int64(dp[0])
}

func main() { bench.Run(67, solve) }
