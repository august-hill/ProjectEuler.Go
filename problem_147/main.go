// Problem 147: Rectangles in Cross-Hatched Grids
// How many rectangles (axis-aligned and diagonal) fit in grids up to 47x43?
// Answer: 846910284

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func solve() int64 {
	var total int64
	for m := 1; m <= 47; m++ {
		for n := 1; n <= 43; n++ {
			// Axis-aligned rectangles
			aa := int64(m) * int64(m+1) / 2 * int64(n) * int64(n+1) / 2

			// Diagonal rectangles
			var diag int64
			for s1 := 1; s1 <= m+n-2; s1++ {
				for s2 := s1 + 1; s2 <= m+n-1; s2++ {
					lo := -s1
					if tmp := s2 - 2*n; tmp > lo {
						lo = tmp
					}
					hi := s1
					if tmp := 2*m - s2; tmp < hi {
						hi = tmp
					}
					if lo >= hi {
						continue
					}
					cnt := int64(hi - lo + 1)
					if cnt >= 2 {
						diag += cnt * (cnt - 1) / 2
					}
				}
			}

			total += aa + diag
		}
	}
	return total
}

func main() { bench.Run(147, solve) }
