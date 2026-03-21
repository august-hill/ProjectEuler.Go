// Answer: 40730
// Problem 34: Digit Factorials
// Find the sum of all numbers which are equal to the sum of the factorial of their digits.

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

var factorials = [10]int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

func digitFactorialSum(n int) int {
	sum := 0
	for n > 0 {
		sum += factorials[n%10]
		n /= 10
	}
	return sum
}

func solve() int64 {
	// Upper bound: 7 * 9! = 2540160
	sum := 0
	for n := 3; n <= 2540160; n++ {
		if n == digitFactorialSum(n) {
			sum += n
		}
	}
	return int64(sum)
}

func main() { bench.Run(34, solve) }
