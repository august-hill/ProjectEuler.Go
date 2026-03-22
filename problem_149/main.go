// Problem 149: Searching for a maximum-sum subsequence
// Find the maximum sum of any sub-sequence in the 2000x2000 grid.
// Answer: 52852124

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const gridN149 = 2000

var table149 [gridN149][gridN149]int64

func initTable149() {
	s := make([]int64, gridN149*gridN149+1)
	for k := 1; k <= 55; k++ {
		s[k] = ((100003 - 200003*int64(k) + 300007*int64(k)*int64(k)*int64(k)) % 1000000 + 1000000) % 1000000 - 500000
	}
	for k := 56; k <= gridN149*gridN149; k++ {
		s[k] = ((s[k-24] + s[k-55] + 1000000) % 1000000 + 1000000) % 1000000 - 500000
	}
	for i := 0; i < gridN149; i++ {
		for j := 0; j < gridN149; j++ {
			table149[i][j] = s[i*gridN149+j+1]
		}
	}
}

func maxSubarray149(arr []int64) int64 {
	best := arr[0]
	current := arr[0]
	for i := 1; i < len(arr); i++ {
		if current < 0 {
			current = arr[i]
		} else {
			current += arr[i]
		}
		if current > best {
			best = current
		}
	}
	return best
}

func solve() int64 {
	initTable149()

	best := int64(-1000000000)
	line := make([]int64, gridN149)

	// Rows
	for i := 0; i < gridN149; i++ {
		val := maxSubarray149(table149[i][:])
		if val > best {
			best = val
		}
	}

	// Columns
	for j := 0; j < gridN149; j++ {
		for i := 0; i < gridN149; i++ {
			line[i] = table149[i][j]
		}
		val := maxSubarray149(line)
		if val > best {
			best = val
		}
	}

	// Diagonals (top-left to bottom-right)
	for start := -(gridN149 - 1); start < gridN149; start++ {
		length := 0
		for i := 0; i < gridN149; i++ {
			j := i - start
			if j >= 0 && j < gridN149 {
				line[length] = table149[i][j]
				length++
			}
		}
		if length > 0 {
			val := maxSubarray149(line[:length])
			if val > best {
				best = val
			}
		}
	}

	// Anti-diagonals
	for start := 0; start < 2*gridN149-1; start++ {
		length := 0
		for i := 0; i < gridN149; i++ {
			j := start - i
			if j >= 0 && j < gridN149 {
				line[length] = table149[i][j]
				length++
			}
		}
		if length > 0 {
			val := maxSubarray149(line[:length])
			if val > best {
				best = val
			}
		}
	}

	return best
}

func main() { bench.Run(149, solve) }
