// Problem 148: Exploring Pascal's Triangle
// How many entries in the first 10^9 rows are not divisible by 7?
// Answer: 2129970655314432

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func solve() int64 {
	N := int64(1000000000) // 10^9

	// Convert N to base 7
	digits := make([]int, 0, 20)
	tmp := N
	for tmp > 0 {
		digits = append(digits, int(tmp%7))
		tmp /= 7
	}
	ndigits := len(digits)
	// digits[0] is LSB

	// Compute 28^i for each position
	pow28 := make([]int64, ndigits)
	pow28[0] = 1
	for i := 1; i < ndigits; i++ {
		pow28[i] = pow28[i-1] * 28
	}

	// f(N) = sum from MSB to LSB
	var total int64
	var multiplier int64 = 1
	for i := ndigits - 1; i >= 0; i-- {
		d := int64(digits[i])
		total += multiplier * d * (d + 1) / 2 * pow28[i]
		multiplier *= d + 1
	}

	return total
}

func main() { bench.Run(148, solve) }
