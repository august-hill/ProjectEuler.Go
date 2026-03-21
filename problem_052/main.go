// Problem 52: Permuted Multiples
// Find smallest x where x, 2x, 3x, 4x, 5x, 6x contain same digits.
// Answer: 142857

package main

import (
	"fmt"
	"sort"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func digitSignature(n int) string {
	s := fmt.Sprintf("%d", n)
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}

func solve() int64 {
	for x := 1; ; x++ {
		sig := digitSignature(x)
		match := true
		for m := 2; m <= 6; m++ {
			if digitSignature(x*m) != sig {
				match = false
				break
			}
		}
		if match {
			return int64(x)
		}
	}
}

func main() { bench.Run(52, solve) }
