// Problem 17: Number Letter Counts
// Count letters used writing 1-1000 in British English.
// Answer: 21124

package main

import (
	"fmt"
	"time"
)

// Letter counts for ones, teens, tens (precomputed for speed)
var ones = []int{0, 3, 3, 5, 4, 4, 3, 5, 5, 4} // "", one, two, three, four, five, six, seven, eight, nine
var teens = []int{3, 6, 6, 8, 8, 7, 7, 9, 8, 8} // ten, eleven, twelve, thirteen, ..., nineteen
var tens = []int{0, 0, 6, 6, 5, 5, 5, 7, 6, 6}  // "", "", twenty, thirty, forty, fifty, sixty, seventy, eighty, ninety

const hundred = 7   // "hundred"
const thousand = 8  // "thousand"
const and = 3       // "and"

func letterCount(n int) int {
	if n == 1000 {
		return ones[1] + thousand // "one thousand"
	}

	count := 0

	// Hundreds place
	if n >= 100 {
		count += ones[n/100] + hundred
		n %= 100
		if n > 0 {
			count += and // British "and"
		}
	}

	// Tens and ones
	if n >= 20 {
		count += tens[n/10]
		count += ones[n%10]
	} else if n >= 10 {
		count += teens[n-10]
	} else {
		count += ones[n]
	}

	return count
}

func solve() int {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += letterCount(i)
	}
	return sum
}

func benchmark(iterations int) time.Duration {
	// Warmup
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
	const iterations = 100000

	fmt.Println("Problem 17: Number Letter Counts")
	fmt.Println("=================================")
	fmt.Printf("Counting letters in 1-1000, Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
