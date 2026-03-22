// Problem 124: Ordered Radicals
// Find E(10000) when all radicals for n <= 100000 are sorted.
// Answer: 21417

package main

import (
	"sort"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

type entry124 struct{ n, rad int }

func solve() int64 {
	const N = 100000
	rad := make([]int, N+1)
	for i := 1; i <= N; i++ {
		rad[i] = 1
	}
	for i := 2; i <= N; i++ {
		if rad[i] == 1 { // i is prime
			for j := i; j <= N; j += i {
				rad[j] *= i
			}
		}
	}

	entries := make([]entry124, N)
	for i := 1; i <= N; i++ {
		entries[i-1] = entry124{i, rad[i]}
	}

	sort.Slice(entries, func(a, b int) bool {
		if entries[a].rad != entries[b].rad {
			return entries[a].rad < entries[b].rad
		}
		return entries[a].n < entries[b].n
	})

	return int64(entries[9999].n)
}

func main() { bench.Run(124, solve) }
