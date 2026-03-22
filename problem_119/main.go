// Problem 119: Digit Power Sum
// Find the 30th term where a_n equals a power of its digit sum.
// Answer: 248155780267521

package main

import (
	"sort"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func digitSum119(n int64) int64 {
	var s int64
	for n > 0 {
		s += n % 10
		n /= 10
	}
	return s
}

func solve() int64 {
	var results []int64

	for base := int64(2); base <= 200; base++ {
		power := base * base
		for exp := 2; exp <= 50 && power < 1e16; exp++ {
			if power >= 10 && digitSum119(power) == base {
				results = append(results, power)
			}
			if power > int64(1e16)/base {
				break
			}
			power *= base
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i] < results[j]
	})

	return results[29]
}

func main() { bench.Run(119, solve) }
