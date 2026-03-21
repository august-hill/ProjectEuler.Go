// Problem 44: Pentagon Numbers
// Find pair of pentagonal numbers where sum and difference are both pentagonal.
// Answer: 5482660

package main

import (
	"sort"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func pentagonal(n int) int {
	return n * (3*n - 1) / 2
}

func solve() int64 {
	const maxN = 3000
	p := make([]int, maxN)
	for i := 0; i < maxN; i++ {
		p[i] = pentagonal(i)
	}

	for j := 1; j < maxN; j++ {
		for k := j + 1; k < maxN; k++ {
			sum := p[j] + p[k]
			if sort.SearchInts(p, sum) >= maxN || p[sort.SearchInts(p, sum)] != sum {
				continue
			}
			diff := p[k] - p[j]
			if sort.SearchInts(p, diff) >= maxN || p[sort.SearchInts(p, diff)] != diff {
				continue
			}
			return int64(diff)
		}
	}
	return 0
}

func main() { bench.Run(44, solve) }
