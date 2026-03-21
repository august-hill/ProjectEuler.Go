// Problem 72: Counting Fractions
// How many reduced proper fractions with d <= 1,000,000?
// Answer: 303963552391

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

const limit072 = 1000001

func solve() int64 {
	phi := make([]int, limit072)
	for i := range phi {
		phi[i] = i
	}

	for i := 2; i < limit072; i++ {
		if phi[i] == i {
			for j := i; j < limit072; j += i {
				phi[j] -= phi[j] / i
			}
		}
	}

	total := int64(0)
	for i := 2; i < limit072; i++ {
		total += int64(phi[i])
	}
	return total
}

func main() { bench.Run(72, solve) }
