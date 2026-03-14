// Problem 44: Pentagon Numbers
// Find pair of pentagonal numbers where sum and difference are both pentagonal.
// Answer: 5482660

package main

import (
	"fmt"
	"sort"
	"time"
)

func pentagonal(n int) int {
	return n * (3*n - 1) / 2
}

func solve() int {
	const maxN = 3000
	p := make([]int, maxN)
	for i := 0; i < maxN; i++ {
		p[i] = pentagonal(i)
	}

	for j := 1; j < maxN; j++ {
		for k := j + 1; k < maxN; k++ {
			sum := p[j] + p[k]
			if sort.SearchInts(p, sum) >= maxN || p[sort.SearchInts(p, sum)] != sum {
				continue
			}
			diff := p[k] - p[j]
			if sort.SearchInts(p, diff) >= maxN || p[sort.SearchInts(p, diff)] != diff {
				continue
			}
			return diff
		}
	}
	return 0
}

func benchmark(iterations int) time.Duration {
	for i := 0; i < 10; i++ {
		solve()
	}
	start := time.Now()
	var result int
	for i := 0; i < iterations; i++ {
		result = solve()
	}
	elapsed := time.Since(start)
	fmt.Printf("Result: %d (%.2f ms/op)\n", result, float64(elapsed.Milliseconds())/float64(iterations))
	return elapsed
}

func main() {
	fmt.Println("Problem 44: Pentagon Numbers")
	fmt.Println("=============================")
	benchmark(100)
}
