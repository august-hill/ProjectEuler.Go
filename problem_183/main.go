// Problem 183: Maximum Product of Parts
// Answer: 48861552

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func gcd183(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func isTerminating183(num, den int64) bool {
	g := gcd183(num, den)
	den /= g
	for den%2 == 0 {
		den /= 2
	}
	for den%5 == 0 {
		den /= 5
	}
	return den == 1
}

func solve() int64 {
	e := math.E
	total := int64(0)

	for N := 5; N <= 10000; N++ {
		kf := int(math.Floor(float64(N) / e))
		kc := kf + 1
		if kf < 2 {
			kf = 2
		}
		vf := float64(kf) * math.Log(float64(N)/float64(kf))
		vc := float64(kc) * math.Log(float64(N)/float64(kc))
		k := kf
		if vc > vf {
			k = kc
		}
		if isTerminating183(int64(N), int64(k)) {
			total -= int64(N)
		} else {
			total += int64(N)
		}
	}
	return total
}

func main() { bench.Run(183, solve) }
