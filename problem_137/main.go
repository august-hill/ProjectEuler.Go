// Problem 137: Fibonacci Golden Nuggets
// Find the 15th golden nugget.
// Answer: 1120149658760

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func solve() int64 {
	fib := make([]int64, 65)
	fib[1] = 1
	fib[2] = 1
	for i := 3; i <= 62; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	// 15th golden nugget = F(30) * F(31)
	return fib[30] * fib[31]
}

func main() { bench.Run(137, solve) }
