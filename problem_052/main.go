// Problem 52: Permuted Multiples
// Find smallest x where x, 2x, 3x, 4x, 5x, 6x contain same digits.
// Answer: 142857

package main

import (
	"fmt"
	"sort"
	"time"
)

func digitSignature(n int) string {
	s := fmt.Sprintf("%d", n)
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}

func solve() int {
	for x := 1; ; x++ {
		sig := digitSignature(x)
		match := true
		for m := 2; m <= 6; m++ {
			if digitSignature(x*m) != sig {
				match = false
				break
			}
		}
		if match {
			return x
		}
	}
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
	fmt.Printf("Result: %d (%.2f ns/op)\n", result, float64(elapsed.Nanoseconds())/float64(iterations))
	return elapsed
}

func main() {
	fmt.Println("Problem 52: Permuted Multiples")
	fmt.Println("===============================")
	benchmark(1000)
}
