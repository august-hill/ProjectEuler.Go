// Answer: 443839
// Problem 30: Digit Fifth Powers
// Find the sum of all numbers that can be written as the sum of fifth powers of their digits.

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

var pow5 = [10]int{0, 1, 32, 243, 1024, 3125, 7776, 16807, 32768, 59049}

func fifthPowerSum(n int) int {
	sum := 0
	for n > 0 {
		sum += pow5[n%10]
		n /= 10
	}
	return sum
}

func solve() int64 {
	// Upper bound: 6 * 9^5 = 354294
	sum := 0
	for n := 2; n <= 354294; n++ {
		if n == fifthPowerSum(n) {
			sum += n
		}
	}
	return int64(sum)
}

func main() { bench.Run(30, solve) }
