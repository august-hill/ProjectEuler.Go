// Problem 158: Exploring strings for which only one character comes lexicographically
// after its neighbour to the left.
// Answer: 409511334375

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	best := int64(0)
	comb := int64(1) // C(26, 0)
	for n := 1; n <= 26; n++ {
		comb = comb * int64(26-n+1) / int64(n)
		pow2 := int64(1) << uint(n)
		euler := pow2 - int64(n) - 1
		pn := comb * euler
		if pn > best {
			best = pn
		}
	}
	return best
}

func main() { bench.Run(158, solve) }
