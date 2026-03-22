// Problem 151: Paper Sheets of Standard Sizes: An Expected-Value Problem
// Answer: 464399 (representing 0.464399)

package main

import (
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

var (
	once151   sync.Once
	memo151   [2][4][8][16]float64
	vis151    [2][4][8][16]bool
)

func expected151(a2, a3, a4, a5 int) float64 {
	total := a2 + a3 + a4 + a5
	if total == 0 {
		return 0.0
	}
	if a2 == 0 && a3 == 0 && a4 == 0 && a5 == 1 {
		return 0.0
	}
	if vis151[a2][a3][a4][a5] {
		return memo151[a2][a3][a4][a5]
	}
	vis151[a2][a3][a4][a5] = true

	single := 0.0
	if total == 1 {
		single = 1.0
	}
	result := single

	if a5 > 0 {
		result += float64(a5) / float64(total) * expected151(a2, a3, a4, a5-1)
	}
	if a4 > 0 {
		result += float64(a4) / float64(total) * expected151(a2, a3, a4-1, a5+1)
	}
	if a3 > 0 {
		result += float64(a3) / float64(total) * expected151(a2, a3-1, a4+1, a5+1)
	}
	if a2 > 0 {
		result += float64(a2) / float64(total) * expected151(a2-1, a3+1, a4+1, a5+1)
	}

	memo151[a2][a3][a4][a5] = result
	return result
}

func solve() int64 {
	once151.Do(func() {})
	result := expected151(1, 1, 1, 1)
	return int64(result*1000000 + 0.5)
}

func main() { bench.Run(151, solve) }
