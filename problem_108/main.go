// Problem 108: Diophantine Reciprocals I
// Find the smallest n with over 1000 solutions to 1/x + 1/y = 1/n.
// Answer: 180180

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

var primes108 = []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43}

var best108 int64

func search108(idx int, n, divCount int64, maxExp int) {
	solutions := (divCount + 1) / 2
	if solutions > 1000 {
		if n < best108 {
			best108 = n
		}
		return
	}
	if idx >= len(primes108) {
		return
	}

	for e := maxExp; e >= 1; e-- {
		newN := n
		overflow := false
		for j := 0; j < e; j++ {
			if newN > best108/primes108[idx] {
				overflow = true
				break
			}
			newN *= primes108[idx]
		}
		if overflow || newN >= best108 {
			continue
		}
		search108(idx+1, newN, divCount*int64(2*e+1), e)
	}
}

func solve() int64 {
	best108 = 1e18
	search108(0, 1, 1, 20)
	return best108
}

func main() { bench.Run(108, solve) }
