// Problem 139: Pythagorean Tiles
// Find how many Pythagorean families allow tiling a right triangle.
// Answer: 10057761

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const perimLimit139 = 100000000

func gcd139(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func solve() int64 {
	var total int64

	for m := int64(2); 2*m*(m+1) < perimLimit139; m++ {
		for n := int64(1); n < m; n++ {
			if (m+n)%2 == 0 {
				continue
			}
			if gcd139(int(m), int(n)) != 1 {
				continue
			}

			a := m*m - n*n
			b := 2 * m * n
			c := m*m + n*n
			perim := a + b + c
			if perim >= perimLimit139 {
				break
			}

			gap := a - b
			if gap < 0 {
				gap = -gap
			}
			if c%gap == 0 {
				total += (perimLimit139 - 1) / perim
			}
		}
	}
	return total
}

func main() { bench.Run(139, solve) }
