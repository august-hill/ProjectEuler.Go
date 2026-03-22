// Problem 156: Counting Digits
// Answer: 21295121502550

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func countDigit156(n int64, d int) int64 {
	if n <= 0 {
		return 0
	}
	count := int64(0)
	factor := int64(1)
	for factor <= n {
		higher := n / (factor * 10)
		curr := (n / factor) % 10
		lower := n % factor
		if curr < int64(d) {
			count += higher * factor
		} else if curr == int64(d) {
			count += higher*factor + lower + 1
		} else {
			count += (higher + 1) * factor
		}
		factor *= 10
	}
	return count
}

var sumFixed156 int64

func findZeros156(d int, lo, hi int64) {
	if lo > hi {
		return
	}
	gLo := countDigit156(lo, d) - lo
	gHi := countDigit156(hi, d) - hi

	if lo == hi {
		if gLo == 0 {
			sumFixed156 += lo
		}
		return
	}

	if gLo > 0 && gHi > 0 {
		if gLo > hi-lo && gHi > hi-lo {
			return
		}
	}
	if gLo < 0 && gHi < 0 {
		if -gLo > 12*(hi-lo) && -gHi > 12*(hi-lo) {
			return
		}
	}

	if hi-lo < 1000 {
		for n := lo; n <= hi; n++ {
			if countDigit156(n, d) == n {
				sumFixed156 += n
			}
		}
		return
	}

	mid := lo + (hi-lo)/2
	findZeros156(d, lo, mid)
	findZeros156(d, mid+1, hi)
}

func solve() int64 {
	sumFixed156 = 0
	for d := 1; d <= 9; d++ {
		findZeros156(d, 1, 200000000000)
	}
	return sumFixed156
}

func main() { bench.Run(156, solve) }
