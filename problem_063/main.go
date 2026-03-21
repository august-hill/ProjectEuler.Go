// Problem 063: Powerful Digit Counts
// How many n-digit positive integers exist which are also an nth power?
// Answer: 49

package main

import (
	"math/big"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	count := 0

	// Base can only be 1-9 (10^n always has n+1 digits)
	for base := 1; base <= 9; base++ {
		power := big.NewInt(int64(base))
		b := big.NewInt(int64(base))
		n := 1

		for {
			digits := len(power.String())

			if digits == n {
				count++
			} else if digits < n {
				break
			}

			power.Mul(power, b)
			n++

			if n > 100 {
				break
			}
		}
	}

	return int64(count)
}

func main() { bench.Run(63, solve) }
