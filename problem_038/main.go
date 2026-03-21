// Problem 38: Pandigital Multiples
// What is the largest 1 to 9 pandigital 9-digit number that can be formed as the
// concatenated product of an integer with (1,2,...,n) where n > 1?
// Answer: 932718654

package main

import (
	"strconv"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func isPandigital(s string) bool {
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
	largest := int64(0)

	for num := 1; num < 10000; num++ {
		concat := ""
		n := 1

		for len(concat) < 9 {
			concat += strconv.Itoa(num * n)
			n++
		}

		if n > 2 && isPandigital(concat) {
			val, _ := strconv.ParseInt(concat, 10, 64)
			if val > largest {
				largest = val
			}
		}
	}

	return largest
}

func main() { bench.Run(38, solve) }
