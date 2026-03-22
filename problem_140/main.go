// Problem 140: Modified Fibonacci Golden Nuggets
// Find the sum of the first 30 golden nuggets for the modified series.
// Answer: 5673835352990

package main

import (
	"sort"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	// Solve m^2 - 5k^2 = 44 with seeds and Pell recurrence (9m+20k, 4m+9k)
	seeds := [][2]int64{
		{7, 1}, {8, 2}, {13, 5}, {17, 7}, {32, 14}, {43, 19},
	}

	nuggets := make([]int64, 0, 200)

	for _, seed := range seeds {
		m, k := seed[0], seed[1]
		for iter := 0; iter < 40 && len(nuggets) < 200; iter++ {
			if m > 7 && m%5 == 2 {
				n := (m - 7) / 5
				if n > 0 {
					nuggets = append(nuggets, n)
				}
			}
			nm := 9*m + 20*k
			nk := 4*m + 9*k
			m, k = nm, nk
		}
	}

	sort.Slice(nuggets, func(i, j int) bool {
		return nuggets[i] < nuggets[j]
	})

	// Remove duplicates and sum first 30
	var sum int64
	unique := 0
	for i := 0; i < len(nuggets) && unique < 30; i++ {
		if i == 0 || nuggets[i] != nuggets[i-1] {
			sum += nuggets[i]
			unique++
		}
	}
	return sum
}

func main() { bench.Run(140, solve) }
