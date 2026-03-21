// Problem 55: Lychrel Numbers
// How many Lychrel numbers are there below ten-thousand?
// Answer: 249

package main

import (
	"math/big"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func reverse(n *big.Int) *big.Int {
	s := n.String()
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	rev := new(big.Int)
	rev.SetString(string(runes), 10)
	return rev
}

func isPalindrome(n *big.Int) bool {
	s := n.String()
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func isLychrel(n int) bool {
	val := big.NewInt(int64(n))
	for i := 0; i < 50; i++ {
		val.Add(val, reverse(val))
		if isPalindrome(val) {
			return false
		}
	}
	return true
}

func solve() int64 {
	count := 0
	for n := 1; n < 10000; n++ {
		if isLychrel(n) {
			count++
		}
	}
	return int64(count)
}

func main() { bench.Run(55, solve) }
