// Problem 166: Criss Cross
// Answer: 7130034

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	count := int64(0)

	for S := 0; S <= 36; S++ {
		for a := 0; a <= 9; a++ {
			for b := 0; b <= 9; b++ {
				for c := 0; c <= 9; c++ {
					d := S - a - b - c
					if d < 0 || d > 9 {
						continue
					}
					for e := 0; e <= 9; e++ {
						for f := 0; f <= 9; f++ {
							for i := 0; i <= 9; i++ {
								mVal := S - a - e - i
								if mVal < 0 || mVal > 9 {
									continue
								}
								for j := 0; j <= 9; j++ {
									n := S - b - f - j
									if n < 0 || n > 9 {
										continue
									}
									g := 2*a + b + c + e + i - j - S
									if g < 0 || g > 9 {
										continue
									}
									h := 2*S - 2*a - b - c - 2*e - f - i + j
									if h < 0 || h > 9 {
										continue
									}
									o := f + j - c
									if o < 0 || o > 9 {
										continue
									}
									k := 2*S - 2*a - b - c - e - f - i
									if k < 0 || k > 9 {
										continue
									}
									l := -S + 2*a + b + c + e + f - j
									if l < 0 || l > 9 {
										continue
									}
									p := a + b + c + e + i - S
									if p < 0 || p > 9 {
										continue
									}
									count++
								}
							}
						}
					}
				}
			}
		}
	}
	return count
}

func main() { bench.Run(166, solve) }
