// Problem 171: Finding numbers for which f(n) is a perfect square
// Answer: 142989277 (last 9 digits)

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const mod171 = 1000000000
const maxSS171 = 1621

func isPerfectSquare171(n int) bool {
	if n < 0 {
		return false
	}
	s := int(math.Sqrt(float64(n)))
	for s*s > n {
		s--
	}
	for (s+1)*(s+1) <= n {
		s++
	}
	return s*s == n
}

func solve() int64 {
	total := int64(0)

	for L := 1; L <= 20; L++ {
		cntDp := make([]int64, maxSS171)
		valDp := make([]int64, maxSS171)

		for d := 1; d <= 9; d++ {
			ss := d * d
			cntDp[ss] = (cntDp[ss] + 1) % mod171
			valDp[ss] = (valDp[ss] + int64(d)) % mod171
		}

		for pos := 1; pos < L; pos++ {
			newCnt := make([]int64, maxSS171)
			newVal := make([]int64, maxSS171)

			for ss := 0; ss < maxSS171; ss++ {
				if cntDp[ss] == 0 {
					continue
				}
				c := cntDp[ss]
				v := valDp[ss]

				for d := 0; d <= 9; d++ {
					nss := ss + d*d
					if nss >= maxSS171 {
						continue
					}
					newCnt[nss] = (newCnt[nss] + c) % mod171
					newVal[nss] = (newVal[nss] + v*10 + c*int64(d)) % mod171
				}
			}
			cntDp = newCnt
			valDp = newVal
		}

		for ss := 0; ss < maxSS171; ss++ {
			if cntDp[ss] > 0 && isPerfectSquare171(ss) {
				total = (total + valDp[ss]) % mod171
			}
		}
	}
	return total
}

func main() { bench.Run(171, solve) }
