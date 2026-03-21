// Problem 36: Double-base Palindromes
// Find the sum of all numbers less than one million which are palindromic in both base 10 and base 2.
// Answer: 872187

package main

import (
	"strconv"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func solve() int64 {
	sum := 0

	for n := 1; n < 1000000; n++ {
		dec := strconv.Itoa(n)
		bin := strconv.FormatInt(int64(n), 2)

		if isPalindrome(dec) && isPalindrome(bin) {
			sum += n
		}
	}

	return int64(sum)
}

func main() { bench.Run(36, solve) }
