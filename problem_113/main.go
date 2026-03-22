// Problem 113: Non-bouncy Numbers Below a Googol
// How many numbers below 10^100 are not bouncy?
// Answer: 51161058134250

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func comb113(n, k int) int64 {
	if k > n || k < 0 {
		return 0
	}
	if k > n-k {
		k = n - k
	}
	// Use int64; values here are large but fit
	var result int64 = 1
	for i := 0; i < k; i++ {
		result = result * int64(n-i) / int64(i+1)
	}
	return result
}

func solve() int64 {
	n := 100
	increasing := comb113(n+9, 9) - 1
	decreasing := comb113(n+10, 10) - int64(n+1)
	flat := int64(9 * n)
	return increasing + decreasing - flat
}

func main() { bench.Run(113, solve) }
