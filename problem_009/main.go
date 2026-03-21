// Answer: 31875000
// Problem 9: Special Pythagorean Triplet
// Find the Pythagorean triplet where a + b + c = 1000.

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const target = 1000

// Optimized: O(n) - derive b algebraically from both constraints
func solve() int64 {
	s := target
	for a := 1; a < s/3; a++ {
		numerator := s*s - 2*s*a
		denominator := 2*s - 2*a
		if numerator%denominator == 0 {
			b := numerator / denominator
			c := s - a - b
			if a < b && b < c && a*a+b*b == c*c {
				return int64(a * b * c)
			}
		}
	}
	return 0
}

func main() { bench.Run(9, solve) }
