// Problem 19: Counting Sundays
// How many Sundays fell on the first of the month during the 20th century?
// Answer: 171

package main

import (
	"fmt"
	"time"
)

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func daysInMonth(month, year int) int {
	days := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if month == 2 && isLeapYear(year) {
		return 29
	}
	return days[month]
}

func solve() int {
	// Jan 1, 1900 was a Monday (day_of_week = 1, where 0 = Sunday)
	dayOfWeek := 1
	count := 0

	// Advance through 1900
	for month := 1; month <= 12; month++ {
		dayOfWeek = (dayOfWeek + daysInMonth(month, 1900)) % 7
	}

	// Count Sundays on the 1st from 1901 to 2000
	for year := 1901; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			if dayOfWeek == 0 {
				count++
			}
			dayOfWeek = (dayOfWeek + daysInMonth(month, year)) % 7
		}
	}

	return count
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
	const iterations = 10000

	fmt.Println("Problem 19: Counting Sundays")
	fmt.Println("============================")
	fmt.Printf("Counting Sundays on 1st of month (1901-2000), Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
