// Answer: 45228
// Problem 32: Pandigital Products
// Find the sum of all products whose multiplicand/multiplier/product identity
// can be written as a 1 through 9 pandigital.

package main

import (
	"fmt"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func isPandigital(a, b, c int) bool {
	s := fmt.Sprintf("%d%d%d", a, b, c)
	if len(s) != 9 {
		return false
	}

	var digits [10]bool
	for _, ch := range s {
		d := int(ch - '0')
		if d == 0 || digits[d] {
			return false
		}
		digits[d] = true
	}
	return true
}

func solve() int64 {
	products := make(map[int]bool)

	for a := 1; a < 100; a++ {
		start := 1000
		end := 9999
		if a >= 10 {
			start = 100
			end = 999
		}

		for b := start; b <= end; b++ {
			c := a * b
			if isPandigital(a, b, c) {
				products[c] = true
			}
		}
	}

	sum := 0
	for p := range products {
		sum += p
	}
	return int64(sum)
}

func main() { bench.Run(32, solve) }
