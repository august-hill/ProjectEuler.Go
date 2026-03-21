// Answer: 2783915460
// Problem 24: Lexicographic Permutations
// What is the millionth lexicographic permutation of the digits 0-9?

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func factorial(n int) int64 {
	result := int64(1)
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}
	return result
}

func solve() int64 {
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := int64(999999) // 0-indexed
	result := int64(0)

	for i := 9; i >= 1; i-- {
		fact := factorial(i)
		idx := int(n / fact)
		result = result*10 + int64(digits[idx])

		// Remove digit at idx
		digits = append(digits[:idx], digits[idx+1:]...)
		n %= fact
	}
	result = result*10 + int64(digits[0])

	return result
}

func main() { bench.Run(24, solve) }
