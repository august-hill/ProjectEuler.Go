// Problem 086: Cuboid Route
// Find least M such that the number of cuboid shortest routes exceeds one million.
// Answer: 1818

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	count := 0
	m := 1

	for {
		for s := 2; s <= 2*m; s++ {
			sq := m*m + s*s
			root := int(math.Sqrt(float64(sq)))
			if root*root == sq {
				cMin := 1
				if s > m {
					cMin = s - m
				}
				cMax := s / 2
				if cMax >= cMin {
					count += cMax - cMin + 1
				}
			}
		}

		if count > 1000000 {
			return int64(m)
		}
		m++
	}
}

func main() { bench.Run(86, solve) }
