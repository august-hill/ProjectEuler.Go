// Problem 110: Diophantine Reciprocals II
// Find the smallest n with over 4,000,000 solutions to 1/x + 1/y = 1/n.
// Answer: 9350130049860600

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

var primes110 = []float64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}

var best110 float64
var bestExps110 [15]int

func search110(idx int, logn float64, divCount int64, maxExp int, exps []int) {
	if divCount > 7999999 {
		if logn < best110 {
			best110 = logn
			copy(bestExps110[:], exps)
		}
		return
	}

	if idx >= len(primes110) {
		return
	}

	// Pruning
	var remaining int64 = 1
	for i := idx; i < len(primes110); i++ {
		remaining *= int64(2*maxExp + 1)
		if remaining > 1e18 {
			break
		}
	}
	if divCount*remaining <= 7999999 {
		return
	}

	for e := maxExp; e >= 1; e-- {
		newLogn := logn + float64(e)*math.Log(primes110[idx])
		if newLogn >= best110 {
			continue
		}
		exps[idx] = e
		search110(idx+1, newLogn, divCount*int64(2*e+1), e, exps)
	}
	exps[idx] = 0
	search110(idx+1, logn, divCount, maxExp, exps)
}

func solve() int64 {
	best110 = 1e30
	exps := make([]int, len(primes110))
	search110(0, 0.0, 1, 7, exps)

	var result int64 = 1
	for i := 0; i < len(primes110); i++ {
		for j := 0; j < bestExps110[i]; j++ {
			result *= int64(primes110[i])
		}
	}
	return result
}

func main() { bench.Run(110, solve) }
