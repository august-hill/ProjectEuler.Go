// Problem 001: Multiples of 3 or 5
// Find the sum of all multiples of 3 or 5 below 1000.

package main

import (
	"fmt"
	"time"
)

// bruteForce iterates through all numbers and checks divisibility
func bruteForce(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

// arithmetic uses inclusion-exclusion with arithmetic series formula
func arithmetic(n int) int {
	return sumMultiples(3, n) + sumMultiples(5, n) - sumMultiples(15, n)
}

// sumMultiples returns sum of all multiples of k below n
func sumMultiples(k, n int) int {
	m := (n - 1) / k
	return k * m * (m + 1) / 2
}

func main() {
	n := 1000

	start1 := time.Now()
	result1 := bruteForce(n)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	result2 := arithmetic(n)
	elapsed2 := time.Since(start2)

	fmt.Printf("Brute Force: %d  (%v)\n", result1, elapsed1)
	fmt.Printf("Arithmetic:  %d  (%v)\n", result2, elapsed2)

	if result1 != result2 {
		fmt.Println("WARNING: Results do not match!")
	}
}
