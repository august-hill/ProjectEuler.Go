// Problem 142: Perfect Square Collection
// Find the smallest x+y+z where x,y,z give 6 perfect squares.
// Answer: 1006193

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func isqrtCheck142(n int64) bool {
	if n < 0 {
		return false
	}
	s := int64(math.Sqrt(float64(n)))
	for s*s < n {
		s++
	}
	for s*s > n {
		s--
	}
	return s*s == n
}

func solve() int64 {
	for a := 3; ; a++ {
		for c := a - 1; c >= 2; c-- {
			f2 := int64(a*a - c*c)
			if !isqrtCheck142(f2) {
				continue
			}
			for e := c - 1; e >= 1; e-- {
				b2 := int64(c*c - e*e)
				if !isqrtCheck142(b2) {
					continue
				}
				d2 := int64(a*a - e*e)
				if !isqrtCheck142(d2) {
					continue
				}

				bv := int64(math.Sqrt(float64(b2)))
				for bv*bv < b2 {
					bv++
				}
				if (int64(a)+bv)%2 != 0 {
					continue
				}

				a2 := int64(a * a)
				x := (a2 + b2) / 2
				y := (a2 - b2) / 2
				z := x - d2

				if z <= 0 || y <= z {
					continue
				}

				return x + y + z
			}
		}
	}
}

func main() { bench.Run(142, solve) }
