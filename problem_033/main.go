// Answer: 100
// Problem 33: Digit Cancelling Fractions
// Find the denominator of the product of the four "curious" fractions in lowest terms.

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func solve() int64 {
	numProduct := 1
	denProduct := 1

	for a := 1; a <= 9; a++ {
		for b := 1; b <= 9; b++ {
			for c := 1; c <= 9; c++ {
				for d := 1; d <= 9; d++ {
					num := a*10 + b
					den := c*10 + d

					// Fraction must be less than 1
					if num >= den {
						continue
					}

					// Check if "cancelling" gives equivalent fraction
					// ab/cd: if b == c, check if a/d == (ab)/(cd)
					if b == c && a*den == num*d {
						numProduct *= num
						denProduct *= den
					}
				}
			}
		}
	}

	return int64(denProduct / gcd(numProduct, denProduct))
}

func main() { bench.Run(33, solve) }
