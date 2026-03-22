// Problem 106: Special Subset Sums: Meta-testing
// How many pairs of subsets of a size-12 set need to be tested?
// Answer: 21384

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func comb106(n, k int) int64 {
	if k > n || k < 0 {
		return 0
	}
	if k > n-k {
		k = n - k
	}
	var result int64 = 1
	for i := 0; i < k; i++ {
		result = result * int64(n-i) / int64(i+1)
	}
	return result
}

func catalan106(n int) int64 {
	return comb106(2*n, n) / int64(n+1)
}

func solve() int64 {
	n := 12
	var total int64
	for k := 2; k <= n/2; k++ {
		pairsTotal := comb106(n, 2*k) * comb106(2*k, k) / 2
		pairsOk := comb106(n, 2*k) * catalan106(k)
		total += pairsTotal - pairsOk
	}
	return total
}

func main() { bench.Run(106, solve) }
