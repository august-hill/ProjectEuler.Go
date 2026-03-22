// Problem 150: Searching a triangular array for a sub-triangle having minimum-sum
// Answer: -271248680

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const rows150 = 1000

func solve() int64 {
	// Build prefix sums per row using LCG
	prefix := make([][]int64, rows150)
	t := int64(0)
	for r := 0; r < rows150; r++ {
		prefix[r] = make([]int64, r+2)
		prefix[r][0] = 0
		for j := 0; j <= r; j++ {
			t = (615949*t + 797807) & ((1 << 20) - 1)
			s := t - (1 << 19)
			prefix[r][j+1] = prefix[r][j] + s
		}
	}

	minSum := int64(0x7FFFFFFFFFFFFFFF)

	for r := 0; r < rows150; r++ {
		for c := 0; c <= r; c++ {
			var sum int64
			for size := 0; r+size < rows150; size++ {
				row := r + size
				sum += prefix[row][c+size+1] - prefix[row][c]
				if sum < minSum {
					minSum = sum
				}
			}
		}
	}

	return minSum
}

func main() { bench.Run(150, solve) }
