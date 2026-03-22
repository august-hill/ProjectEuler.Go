// Problem 177: Integer Angled Quadrilaterals
// Answer: 129325

package main

import (
	"math"
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

var (
	once177    sync.Once
	sinval177  [181]float64
	result177  int64
)

func init177() {
	for i := 0; i <= 180; i++ {
		sinval177[i] = math.Sin(float64(i) * math.Pi / 180.0)
	}
}

func compute177() {
	once177.Do(init177)
	count := 0

	for P := 1; P <= 179; P++ {
		S := 180 - P
		for a2 := 1; a2 < S; a2++ {
			b1 := S - a2
			for b2 := 1; b2 < P; b2++ {
				c1 := P - b2
				for c2 := 1; c2 < S; c2++ {
					d1 := S - c2
					lhsPartial := sinval177[a2] * sinval177[b2] * sinval177[c2]
					rhsPartial := sinval177[b1] * sinval177[c1] * sinval177[d1]

					for a1 := 1; a1 < P; a1++ {
						d2 := P - a1
						if a1+a2 >= 180 {
							continue
						}
						if b1+b2 >= 180 {
							continue
						}
						if c1+c2 >= 180 {
							continue
						}
						if d1+d2 >= 180 {
							continue
						}

						lhs := lhsPartial * sinval177[d2]
						rhs := rhsPartial * sinval177[a1]

						if math.Abs(lhs-rhs) < 1e-8*(lhs+rhs+1e-15) {
							// Check canonical form
							reps := [8][8]int{
								{a1, a2, b1, b2, c1, c2, d1, d2},
								{b1, b2, c1, c2, d1, d2, a1, a2},
								{c1, c2, d1, d2, a1, a2, b1, b2},
								{d1, d2, a1, a2, b1, b2, c1, c2},
								{a2, a1, d2, d1, c2, c1, b2, b1},
								{d2, d1, c2, c1, b2, b1, a2, a1},
								{c2, c1, b2, b1, a2, a1, d2, d1},
								{b2, b1, a2, a1, d2, d1, c2, c1},
							}
							minIdx := 0
							for r := 1; r < 8; r++ {
								for j := 0; j < 8; j++ {
									if reps[r][j] < reps[minIdx][j] {
										minIdx = r
										break
									}
									if reps[r][j] > reps[minIdx][j] {
										break
									}
								}
							}
							if minIdx == 0 {
								count++
							}
						}
					}
				}
			}
		}
	}
	result177 = int64(count)
}

func solve() int64 {
	once177.Do(compute177)
	return result177
}

func main() { bench.Run(177, solve) }
