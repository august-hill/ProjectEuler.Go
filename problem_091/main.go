// Problem 091: Right Triangles in Quadrants
// How many right triangles with vertices at O(0,0), P(x1,y1), Q(x2,y2)
// where 0 <= x1,y1,x2,y2 <= 50?
// Answer: 14234

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	const n = 50
	count := 0

	for x1 := 0; x1 <= n; x1++ {
		for y1 := 0; y1 <= n; y1++ {
			if x1 == 0 && y1 == 0 {
				continue
			}
			for x2 := 0; x2 <= n; x2++ {
				for y2 := 0; y2 <= n; y2++ {
					if x2 == 0 && y2 == 0 {
						continue
					}
					if x1 == x2 && y1 == y2 {
						continue
					}
					// Avoid double counting
					if x1 > x2 || (x1 == x2 && y1 > y2) {
						continue
					}

					dotO := x1*x2 + y1*y2
					dotP := (-x1)*(x2-x1) + (-y1)*(y2-y1)
					dotQ := (-x2)*(x1-x2) + (-y2)*(y1-y2)

					if dotO == 0 || dotP == 0 || dotQ == 0 {
						count++
					}
				}
			}
		}
	}

	return int64(count)
}

func main() { bench.Run(91, solve) }
