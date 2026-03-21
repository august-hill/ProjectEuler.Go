// Problem 062: Cubic Permutations
// Find the smallest cube for which exactly five permutations of its digits are also cube.
// Answer: 127035954683

package main

import (
	"sort"
	"strconv"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func digitSignature(n int64) string {
	s := strconv.FormatInt(n, 10)
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}

func solve() int64 {
	type entry struct {
		firstCube int64
		count     int
	}
	sigMap := make(map[string]*entry)

	for n := int64(1); n < 100000; n++ {
		cube := n * n * n
		sig := digitSignature(cube)

		e, exists := sigMap[sig]
		if !exists {
			sigMap[sig] = &entry{firstCube: cube, count: 1}
		} else {
			e.count++
			if e.count == 5 {
				return e.firstCube
			}
		}
	}
	return 0
}

func main() { bench.Run(62, solve) }
