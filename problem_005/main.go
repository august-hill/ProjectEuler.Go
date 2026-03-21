// Answer: 232792560
// Problem 005: Smallest Multiple
// Find the smallest positive number divisible by all numbers from 1 to 20.

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

// gcd computes greatest common divisor using Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// lcm computes least common multiple
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

// iterativeLCM computes LCM of 1..n iteratively
func solve() int64 {
	result := 1
	for i := 2; i <= 20; i++ {
		result = lcm(result, i)
	}
	return int64(result)
}

func main() { bench.Run(5, solve) }
