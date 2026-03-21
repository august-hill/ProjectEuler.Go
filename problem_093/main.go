// Problem 093: Arithmetic Expressions
// Find the set of four digits {a,b,c,d} that generates the longest
// set of consecutive positive integers using +, -, *, / and parentheses.
// Answer: 1258

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const maxVal093 = 3500

func evalAll093(a, b float64) []float64 {
	results := []float64{a + b, a - b, b - a, a * b}
	if math.Abs(b) > 1e-12 {
		results = append(results, a/b)
	}
	if math.Abs(a) > 1e-12 {
		results = append(results, b/a)
	}
	return results
}

func solve() int64 {
	bestDigits := 0
	bestCount := 0

	for a := 1; a <= 9; a++ {
		for b := a + 1; b <= 9; b++ {
			for c := b + 1; c <= 9; c++ {
				for d := c + 1; d <= 9; d++ {
					digits := [4]float64{float64(a), float64(b), float64(c), float64(d)}
					var seen [maxVal093]bool

					// All 24 permutations
					perm := [24][4]int{}
					np := 0
					for i := 0; i < 4; i++ {
						for j := 0; j < 4; j++ {
							if j == i {
								continue
							}
							for k := 0; k < 4; k++ {
								if k == i || k == j {
									continue
								}
								l := 6 - i - j - k
								perm[np] = [4]int{i, j, k, l}
								np++
							}
						}
					}

					for p := 0; p < 24; p++ {
						w := digits[perm[p][0]]
						x := digits[perm[p][1]]
						y := digits[perm[p][2]]
						z := digits[perm[p][3]]

						// ((w op x) op y) op z
						wx := evalAll093(w, x)
						for _, wxi := range wx {
							wxy := evalAll093(wxi, y)
							for _, wxyi := range wxy {
								wxyz := evalAll093(wxyi, z)
								for _, r := range wxyz {
									if r > 0.5 && r < maxVal093 && math.Abs(r-math.Round(r)) < 1e-9 {
										seen[int(math.Round(r))] = true
									}
								}
							}
						}

						// (w op x) op (y op z)
						yz := evalAll093(y, z)
						for _, wxi := range wx {
							for _, yzi := range yz {
								res := evalAll093(wxi, yzi)
								for _, r := range res {
									if r > 0.5 && r < maxVal093 && math.Abs(r-math.Round(r)) < 1e-9 {
										seen[int(math.Round(r))] = true
									}
								}
							}
						}
					}

					count := 0
					for n := 1; n < maxVal093; n++ {
						if seen[n] {
							count = n
						} else {
							break
						}
					}
					if count > bestCount {
						bestCount = count
						bestDigits = a*1000 + b*100 + c*10 + d
					}
				}
			}
		}
	}
	return int64(bestDigits)
}

func main() { bench.Run(93, solve) }
