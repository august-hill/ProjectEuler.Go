// Problem 38: Pandigital Multiples
// What is the largest 1 to 9 pandigital 9-digit number that can be formed as the
// concatenated product of an integer with (1,2,...,n) where n > 1?
// Answer: 932718654

package main

import (
	"fmt"
	"strconv"
	"time"
)

func isPandigital(s string) bool {
	if len(s) != 9 {
		return false
	}
	var digits [10]bool
	for _, ch := range s {
		d := int(ch - '0')
		if d == 0 || digits[d] {
			return false
		}
		digits[d] = true
	}
	return true
}

func solve() int64 {
	largest := int64(0)

	for num := 1; num < 10000; num++ {
		concat := ""
		n := 1

		for len(concat) < 9 {
			concat += strconv.Itoa(num * n)
			n++
		}

		if n > 2 && isPandigital(concat) {
			val, _ := strconv.ParseInt(concat, 10, 64)
			if val > largest {
				largest = val
			}
		}
	}

	return largest
}

func benchmark(iterations int) time.Duration {
	// Warmup
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
	const iterations = 1000

	fmt.Println("Problem 38: Pandigital Multiples")
	fmt.Println("=================================")
	fmt.Printf("Finding largest pandigital multiple, Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
