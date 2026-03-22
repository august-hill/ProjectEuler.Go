// Problem 127: abc-hits
// Find the sum of c for all abc-hits with c < 120000.
// Answer: 18407904

package main

import (
	"sort"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const limit127 = 120000

func gcd127(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func solve() int64 {
	rad := make([]int, limit127)
	for i := range rad {
		rad[i] = 1
	}
	for p := 2; p < limit127; p++ {
		if rad[p] == 1 { // p is prime
			for j := p; j < limit127; j += p {
				rad[j] *= p
			}
		}
	}

	sortedByRad := make([]int, limit127)
	for i := range sortedByRad {
		sortedByRad[i] = i
	}
	sort.Slice(sortedByRad, func(i, j int) bool {
		ia, ib := sortedByRad[i], sortedByRad[j]
		if rad[ia] != rad[ib] {
			return rad[ia] < rad[ib]
		}
		return ia < ib
	})

	var total int64
	for c := 2; c < limit127; c++ {
		radLimit := c / rad[c]
		for _, a := range sortedByRad[1:] {
			if rad[a] >= radLimit {
				break
			}
			if a >= c {
				continue
			}
			b := c - a
			if b <= a {
				continue
			}
			if int64(rad[a])*int64(rad[b]) >= int64(radLimit) {
				continue
			}
			if gcd127(a, b) != 1 {
				continue
			}
			total += int64(c)
		}
	}
	return total
}

func main() { bench.Run(127, solve) }
