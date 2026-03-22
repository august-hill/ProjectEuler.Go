// Problem 141: Square Progressive Numbers
// Find sum of all square progressive numbers below 10^12.
// Answer: 878454337159

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const limit141 = 1000000000000

func gcd141(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func solve() int64 {
	var sum int64

	for b := int64(2); b*b*b < limit141; b++ {
		for a := int64(1); a < b; a++ {
			if gcd141(int(a), int(b)) != 1 {
				continue
			}
			for c := int64(1); ; c++ {
				n := a * c * (b*b*b*c + a)
				if n >= limit141 {
					break
				}
				s := int64(math.Sqrt(float64(n)))
				for s*s < n {
					s++
				}
				for s*s > n {
					s--
				}
				if s*s == n {
					sum += n
				}
			}
		}
	}
	return sum
}

func main() { bench.Run(141, solve) }
