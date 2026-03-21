// Answer: 171
// Problem 19: Counting Sundays
// How many Sundays fell on the first of the month during the 20th century?

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

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

func solve() int64 {
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

	return int64(count)
}

func main() { bench.Run(19, solve) }
