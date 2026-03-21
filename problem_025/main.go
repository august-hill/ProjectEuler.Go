// Answer: 4782
// Problem 25: 1000-digit Fibonacci Number
// What is the index of the first term in the Fibonacci sequence to contain 1000 digits?

package main

import (
	"math/big"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	fib1 := big.NewInt(1) // F(1)
	fib2 := big.NewInt(1) // F(2)
	term := 2

	// 10^999 is the smallest 1000-digit number
	threshold := new(big.Int).Exp(big.NewInt(10), big.NewInt(999), nil)

	for fib2.Cmp(threshold) < 0 {
		fib1.Add(fib1, fib2)
		fib1, fib2 = fib2, fib1
		term++
	}

	return int64(term)
}

func main() { bench.Run(25, solve) }
