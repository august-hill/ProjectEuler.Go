// Problem 105: Special Subset Sums: Testing
// Find the sum of elements in all special subset sum sets in the file.
// Answer: 73702

package main

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

//go:embed p105_sets.txt
var setsData string

func isSpecial105(set []int) bool {
	n := len(set)
	limit := 1 << n
	sums := make([]int, limit)
	sizes := make([]int, limit)

	for mask := 0; mask < limit; mask++ {
		s, sz := 0, 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				s += set[i]
				sz++
			}
		}
		sums[mask] = s
		sizes[mask] = sz
	}

	for a := 1; a < limit; a++ {
		for b := a + 1; b < limit; b++ {
			if a&b != 0 {
				continue
			}
			if sums[a] == sums[b] {
				return false
			}
			if sizes[a] > sizes[b] && sums[a] <= sums[b] {
				return false
			}
			if sizes[b] > sizes[a] && sums[b] <= sums[a] {
				return false
			}
		}
	}
	return true
}

func solve() int64 {
	var total int64
	for _, line := range strings.Split(strings.TrimSpace(setsData), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		set := make([]int, 0, len(parts))
		for _, p := range parts {
			v, _ := strconv.Atoi(strings.TrimSpace(p))
			set = append(set, v)
		}
		sort.Ints(set)
		if isSpecial105(set) {
			for _, v := range set {
				total += int64(v)
			}
		}
	}
	return total
}

func main() { bench.Run(105, solve) }
