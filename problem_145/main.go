// Problem 145: How many reversible numbers are there below one-billion?
// Answer: 608720

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func allOdd145(n int64) bool {
	for n > 0 {
		if (n%10)%2 == 0 {
			return false
		}
		n /= 10
	}
	return true
}

func reverseNum145(n int64) int64 {
	var rev int64
	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	return rev
}

func solve() int64 {
	var count int64

	// d=2 through d=7: brute force
	for n := int64(10); n < 10000000; n++ {
		if n%10 == 0 {
			continue
		}
		if allOdd145(n + reverseNum145(n)) {
			count++
		}
	}

	// d=8: 0
	// d=9: 540000
	count += 540000

	return count
}

func main() { bench.Run(145, solve) }
