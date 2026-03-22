// Problem 115: Counting Block Combinations II
// Find the least n such that f(50, n) exceeds one million.
// Answer: 168

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func countWays115(m, n int) int64 {
	temp := make([]int64, n+1)
	temp[0] = 1
	for i := 1; i <= n; i++ {
		temp[i] = temp[i-1]
		for length := m; length <= i; length++ {
			start := i - length
			if start == 0 {
				temp[i]++
			} else if start == 1 {
				temp[i]++
			} else {
				temp[i] += temp[start-1]
			}
		}
	}
	return temp[n]
}

func solve() int64 {
	m := 50
	for n := m; n <= 1000; n++ {
		if countWays115(m, n) > 1000000 {
			return int64(n)
		}
	}
	return -1
}

func main() { bench.Run(115, solve) }
