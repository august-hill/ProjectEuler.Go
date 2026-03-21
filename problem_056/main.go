// Problem 56: Powerful Digit Sum
// Considering natural numbers of the form a^b, where a, b < 100, find the maximum digital sum.
// Answer: 972

package main

import (
	"math/big"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func digitSum(n *big.Int) int {
	sum := 0
	for _, ch := range n.String() {
		sum += int(ch - '0')
	}
	return sum
}

func solve() int64 {
	maxSum := 0
	for a := 2; a < 100; a++ {
		base := big.NewInt(int64(a))
		power := big.NewInt(1)
		for b := 1; b < 100; b++ {
			power.Mul(power, base)
			s := digitSum(power)
			if s > maxSum {
				maxSum = s
			}
		}
	}
	return int64(maxSum)
}

func main() { bench.Run(56, solve) }
