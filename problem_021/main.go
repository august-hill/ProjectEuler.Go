// Answer: 31626
// Problem 21: Amicable Numbers
// Evaluate the sum of all amicable numbers under 10000.

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

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
	sum := 0
	for a := 2; a < 10000; a++ {
		b := sumProperDivisors(a)
		if b != a && b < 10000 && sumProperDivisors(b) == a {
			sum += a
		}
	}
	return int64(sum)
}

func main() { bench.Run(21, solve) }
