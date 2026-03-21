// Problem 80: Square Root Digital Expansion
// For the first 100 natural numbers, find the total of the digital sums
// of the first 100 decimal digits of all irrational square roots.
// Answer: 40886

package main

import (
	"math/big"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func isPerfectSquare080(n int) bool {
	s := 0
	for s*s < n {
		s++
	}
	return s*s == n
}

func sqrtDigitSum080(n int) int {
	// Compute sqrt(n) to 100 decimal digits using big.Int
	// Multiply n by 10^200 (to get 100 digits after decimal point), then take integer sqrt
	multiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(200), nil)
	scaled := new(big.Int).Mul(big.NewInt(int64(n)), multiplier)
	root := new(big.Int).Sqrt(scaled)

	// Sum the first 100 digits
	digits := root.String()
	sum := 0
	count := 0
	for _, ch := range digits {
		if count >= 100 {
			break
		}
		sum += int(ch - '0')
		count++
	}
	return sum
}

func solve() int64 {
	total := 0
	for n := 1; n <= 100; n++ {
		if !isPerfectSquare080(n) {
			total += sqrtDigitSum080(n)
		}
	}
	return int64(total)
}

func main() { bench.Run(80, solve) }
