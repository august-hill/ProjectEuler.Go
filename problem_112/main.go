// Problem 112: Bouncy Numbers
// Find the least number for which the proportion of bouncy numbers is exactly 99%.
// Answer: 1587000

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func isBouncy112(n int) bool {
	increasing, decreasing := false, false
	prev := n % 10
	n /= 10
	for n > 0 {
		d := n % 10
		if d < prev {
			increasing = true
		}
		if d > prev {
			decreasing = true
		}
		if increasing && decreasing {
			return true
		}
		prev = d
		n /= 10
	}
	return false
}

func solve() int64 {
	bouncy := 0
	for n := 1; ; n++ {
		if isBouncy112(n) {
			bouncy++
		}
		if bouncy*100 == n*99 {
			return int64(n)
		}
	}
}

func main() { bench.Run(112, solve) }
