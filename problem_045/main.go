// Problem 45: Triangular, Pentagonal, and Hexagonal
// Find next number after 40755 that is triangular, pentagonal, and hexagonal.
// Answer: 1533776805

package main

import (
	"fmt"
	"time"
)

func pentagonal(n int64) int64 { return n * (3*n - 1) / 2 }
func hexagonal(n int64) int64  { return n * (2*n - 1) }

// Note: Every hexagonal number is also triangular: H(n) = T(2n-1)
// So we just need to find where pentagonal and hexagonal meet.
func solve() int64 {
	p, h := int64(1), int64(1)
	pn, hn := pentagonal(p), hexagonal(h)

	for {
		if hn < pn {
			h++
			hn = hexagonal(h)
		} else if pn < hn {
			p++
			pn = pentagonal(p)
		} else {
			if hn > 40755 {
				return hn
			}
			p++
			h++
			pn = pentagonal(p)
			hn = hexagonal(h)
		}
	}
}

func benchmark(iterations int) time.Duration {
	for i := 0; i < 10; i++ {
		solve()
	}
	start := time.Now()
	var result int64
	for i := 0; i < iterations; i++ {
		result = solve()
	}
	elapsed := time.Since(start)
	fmt.Printf("Result: %d (%.2f ns/op)\n", result, float64(elapsed.Nanoseconds())/float64(iterations))
	return elapsed
}

func main() {
	fmt.Println("Problem 45: Triangular, Pentagonal, Hexagonal")
	fmt.Println("==============================================")
	benchmark(100000)
}
