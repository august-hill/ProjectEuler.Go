// Answer: 233168
// Problem 001: Multiples of 3 or 5
// Find the sum of all multiples of 3 or 5 below 1000.

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

// sumMultiples returns sum of all multiples of k below n
func sumMultiples(k, n int) int {
	m := (n - 1) / k
	return k * m * (m + 1) / 2
}

// arithmetic uses inclusion-exclusion with arithmetic series formula
func solve() int64 {
	n := 1000
	return int64(sumMultiples(3, n) + sumMultiples(5, n) - sumMultiples(15, n))
}

func main() { bench.Run(1, solve) }
