// Answer: 4179871
// Problem 23: Non-abundant Sums
// Find the sum of all positive integers which cannot be written as the sum of two abundant numbers.

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const limit = 28123

func sumProperDivisors(n int) int {
	if n <= 1 {
		return 0
	}
	sum := 1
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			sum += i
			other := n / i
			if other != i {
				sum += other
			}
		}
	}
	return sum
}

func solve() int64 {
	// Find all abundant numbers
	abundants := make([]int, 0, 10000)
	for i := 12; i <= limit; i++ {
		if sumProperDivisors(i) > i {
			abundants = append(abundants, i)
		}
	}

	// Mark all numbers that can be expressed as sum of two abundant numbers
	expressible := make([]bool, limit+1)
	for i := 0; i < len(abundants); i++ {
		for j := i; j < len(abundants); j++ {
			sum := abundants[i] + abundants[j]
			if sum <= limit {
				expressible[sum] = true
			} else {
				break
			}
		}
	}

	// Sum all numbers that cannot be expressed
	result := 0
	for i := 1; i <= limit; i++ {
		if !expressible[i] {
			result += i
		}
	}
	return int64(result)
}

func main() { bench.Run(23, solve) }
