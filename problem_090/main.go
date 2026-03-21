// Problem 090: Cube Digit Pairs
// How many distinct arrangements of two cubes allow all square numbers to be displayed?
// Answer: 1217

package main

import (
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

type cube090 [10]bool

var (
	cubes090    []cube090
	initOnce090 sync.Once
)

func generateCubes090() {
	initOnce090.Do(func() {
		cubes090 = nil
		for a := 0; a < 5; a++ {
			for b := a + 1; b < 6; b++ {
				for c := b + 1; c < 7; c++ {
					for d := c + 1; d < 8; d++ {
						for e := d + 1; e < 9; e++ {
							for f := e + 1; f < 10; f++ {
								var cb cube090
								cb[a] = true
								cb[b] = true
								cb[c] = true
								cb[d] = true
								cb[e] = true
								cb[f] = true
								// 6 and 9 are interchangeable
								if cb[6] || cb[9] {
									cb[6] = true
									cb[9] = true
								}
								cubes090 = append(cubes090, cb)
							}
						}
					}
				}
			}
		}
	})
}

func solve() int64 {
	generateCubes090()

	squares := [][2]int{
		{0, 1}, {0, 4}, {0, 9}, {1, 6}, {2, 5},
		{3, 6}, {4, 9}, {6, 4}, {8, 1},
	}

	count := 0
	for i := 0; i < len(cubes090); i++ {
		for j := i; j < len(cubes090); j++ {
			valid := true
			for _, sq := range squares {
				d1, d2 := sq[0], sq[1]
				if !((cubes090[i][d1] && cubes090[j][d2]) || (cubes090[i][d2] && cubes090[j][d1])) {
					valid = false
					break
				}
			}
			if valid {
				count++
			}
		}
	}
	return int64(count)
}

func main() { bench.Run(90, solve) }
