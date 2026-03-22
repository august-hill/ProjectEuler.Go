// Problem 101: Optimum Polynomial
// The generating function is u(n) = 1 - n + n^2 - n^3 + ... + n^10.
// Use Lagrange interpolation to find the sum of FITs.
// Answer: 37076114526

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func u101(n int64) int64 {
	var val, power int64 = 0, 1
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			val += power
		} else {
			val -= power
		}
		power *= n
	}
	return val
}

func lagrangeEval(y []int64, k int, x int) int64 {
	result := 0.0
	for i := 0; i < k; i++ {
		li := 1.0
		for j := 0; j < k; j++ {
			if j != i {
				li *= float64(x-(j+1)) / float64((i+1)-(j+1))
			}
		}
		result += float64(y[i]) * li
	}
	if result > 0 {
		return int64(result + 0.5)
	}
	return int64(result - 0.5)
}

func solve() int64 {
	y := make([]int64, 11)
	for i := 0; i < 11; i++ {
		y[i] = u101(int64(i + 1))
	}

	var total int64
	for k := 1; k <= 10; k++ {
		predicted := lagrangeEval(y, k, k+1)
		actual := u101(int64(k + 1))
		if predicted != actual {
			total += predicted
		}
	}
	return total
}

func main() { bench.Run(101, solve) }
