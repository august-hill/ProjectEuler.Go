// Problem 176: Right-angled Triangles that Share a Cathetus
// Answer: 96818198400000

package main

import (
	"math"
	"sort"
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

var (
	once176        sync.Once
	answerCache176 int64
)

var smallPrimes176 = []int{3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43}

var (
	bestLog176  float64
	bestExps176 []int
	bestA176    int
)

func cmpDesc176(a, b int) bool { return a > b }

func computeLog176(a int, exps []int) float64 {
	v := float64(a) * math.Log(2.0)
	for i, e := range exps {
		v += float64(e) * math.Log(float64(smallPrimes176[i]))
	}
	return v
}

func findFactorizations176(n, minVal int, factors []int, aExp int) {
	if n >= minVal && n%2 == 1 {
		factors = append(factors, n)
		exps := make([]int, len(factors))
		for i, f := range factors {
			exps[i] = (f - 1) / 2
		}
		sort.Slice(exps, func(i, j int) bool { return exps[i] > exps[j] })
		v := computeLog176(aExp, exps)
		if v < bestLog176 {
			bestLog176 = v
			bestA176 = aExp
			bestExps176 = make([]int, len(exps))
			copy(bestExps176, exps)
		}
	}

	for f := minVal; f*f <= n; f++ {
		if f%2 == 0 {
			continue
		}
		if n%f == 0 {
			findFactorizations176(n/f, f, append(factors, f), aExp)
		}
	}
}

func compute176() {
	bestLog176 = 1e30
	bestExps176 = nil

	// Case 1: odd n, product(2*bi+1) = 95095
	findFactorizations176(95095, 3, nil, 0)

	// Case 2: even n, (2*a-1)*product(2*bi+1) = 95095
	var divs []int
	for d := 1; d*d <= 95095; d++ {
		if 95095%d == 0 {
			divs = append(divs, d)
			if d != 95095/d {
				divs = append(divs, 95095/d)
			}
		}
	}

	for _, d := range divs {
		if d%2 == 0 {
			continue
		}
		a := (d + 1) / 2
		if a < 1 {
			continue
		}
		remaining := 95095 / d
		if remaining == 1 {
			v := float64(a) * math.Log(2.0)
			if v < bestLog176 {
				bestLog176 = v
				bestA176 = a
				bestExps176 = nil
			}
		} else {
			findFactorizations176(remaining, 3, nil, a)
		}
	}

	result := int64(1)
	for i := 0; i < bestA176; i++ {
		result *= 2
	}
	for i, e := range bestExps176 {
		base := int64(smallPrimes176[i])
		for j := 0; j < e; j++ {
			result *= base
		}
	}
	answerCache176 = result
}

func solve() int64 {
	once176.Do(compute176)
	return answerCache176
}

func main() { bench.Run(176, solve) }
