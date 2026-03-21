// Problem 040: Champernowne's Constant
// Find d1 x d10 x d100 x d1000 x d10000 x d100000 x d1000000 of Champernowne's constant.
// Answer: 210

package main

import (
	"strconv"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	champernowne := make([]byte, 0, 1100000)

	for i := 1; len(champernowne) < 1000001; i++ {
		champernowne = append(champernowne, []byte(strconv.Itoa(i))...)
	}

	d := func(pos int) int {
		return int(champernowne[pos-1] - '0')
	}

	return int64(d(1) * d(10) * d(100) * d(1000) * d(10000) * d(100000) * d(1000000))
}

func main() { bench.Run(40, solve) }
