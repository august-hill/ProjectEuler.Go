// Problem 75: Singular Integer Right Triangles
// For how many values of L <= 1,500,000 is there exactly one integer-sided right triangle?
// Answer: 161667

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const limit075 = 1500000

func gcd075(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func solve() int64 {
	countArr := make([]uint8, limit075+1)

	mMax := int(math.Sqrt(float64(limit075 / 2)))

	for m := 2; m <= mMax; m++ {
		for n := 1; n < m; n++ {
			if (m-n)%2 == 0 {
				continue
			}
			if gcd075(m, n) != 1 {
				continue
			}

			perim := 2 * m * (m + n)
			if perim > limit075 {
				break
			}

			for k := perim; k <= limit075; k += perim {
				if countArr[k] < 2 {
					countArr[k]++
				}
			}
		}
	}

	result := 0
	for i := 1; i <= limit075; i++ {
		if countArr[i] == 1 {
			result++
		}
	}

	return int64(result)
}

func main() { bench.Run(75, solve) }
