// Answer: 648
// Problem 20: Factorial Digit Sum
// Find the sum of the digits in 100!

package main

import (
	"math/big"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	factorial := big.NewInt(1)
	for n := 2; n <= 100; n++ {
		factorial.Mul(factorial, big.NewInt(int64(n)))
	}

	sum := 0
	for _, ch := range factorial.String() {
		sum += int(ch - '0')
	}
	return int64(sum)
}

func main() { bench.Run(20, solve) }
