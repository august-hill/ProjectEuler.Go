// Answer: 1366
// Problem 16: Power Digit Sum
// What is the sum of the digits of 2^1000?

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

// Manual digit doubling - no big integer library needed
func powerDigitSum(n int) int {
	digits := []int{1}

	for i := 0; i < n; i++ {
		carry := 0
		for j := 0; j < len(digits); j++ {
			val := digits[j]*2 + carry
			digits[j] = val % 10
			carry = val / 10
		}
		if carry > 0 {
			digits = append(digits, carry)
		}
	}

	sum := 0
	for _, d := range digits {
		sum += d
	}
	return sum
}

func solve() int64 {
	return int64(powerDigitSum(1000))
}

func main() { bench.Run(16, solve) }
