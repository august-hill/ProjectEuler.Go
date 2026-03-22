// Problem 199: Iterative Circle Packing
// Answer: 396087 (representing 0.00396087)

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

var totalArea199 float64

func fillGap199(k1, k2, k3 float64, depth int) {
	if depth == 0 {
		return
	}
	k4 := k1 + k2 + k3 + 2.0*math.Sqrt(k1*k2+k2*k3+k1*k3)
	r4 := 1.0 / k4
	totalArea199 += math.Pi * r4 * r4
	fillGap199(k1, k2, k4, depth-1)
	fillGap199(k1, k3, k4, depth-1)
	fillGap199(k2, k3, k4, depth-1)
}

func solve() int64 {
	R := 1.0
	r := R / (1.0 + 2.0/math.Sqrt(3.0))

	kOuter := -1.0 / R
	kInner := 1.0 / r

	outerArea := math.Pi * R * R
	totalArea199 = 3.0 * math.Pi * r * r

	fillGap199(kInner, kInner, kOuter, 10)
	fillGap199(kInner, kInner, kOuter, 10)
	fillGap199(kInner, kInner, kOuter, 10)
	fillGap199(kInner, kInner, kInner, 10)

	fraction := (outerArea - totalArea199) / outerArea
	return int64(math.Round(fraction * 1e8))
}

func main() { bench.Run(199, solve) }
