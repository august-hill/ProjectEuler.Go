// Problem 104: Pandigital Fibonacci Ends
// Find the first Fibonacci number with pandigital first and last 9 digits.
// Answer: 329468

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func isPandigital104(n int64) bool {
	digits := 0
	for i := 0; i < 9; i++ {
		d := n % 10
		n /= 10
		if d == 0 {
			return false
		}
		if digits&(1<<d) != 0 {
			return false
		}
		digits |= 1 << d
	}
	return digits == 0x3FE
}

func solve() int64 {
	const MOD = 1000000000
	var a, b int64 = 1, 1
	logPhi := math.Log10((1.0 + math.Sqrt(5.0)) / 2.0)
	logSqrt5 := math.Log10(5.0) / 2.0

	for k := 3; ; k++ {
		c := (a + b) % MOD
		a = b
		b = c

		if isPandigital104(b) {
			logFk := float64(k)*logPhi - logSqrt5
			frac := logFk - math.Floor(logFk)
			first9 := int64(math.Pow(10.0, frac+8.0))
			if isPandigital104(first9) {
				return int64(k)
			}
		}
	}
}

func main() { bench.Run(104, solve) }
