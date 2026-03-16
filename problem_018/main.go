// Problem 18: Maximum Path Sum I
// Find the maximum total from top to bottom of the triangle.
// Answer: 1074

package main

import (
	"fmt"
	"time"
)

const rows = 15

var triangle = [rows][rows]int{
	{75},
	{95, 64},
	{17, 47, 82},
	{18, 35, 87, 10},
	{20, 4, 82, 47, 65},
	{19, 1, 23, 75, 3, 34},
	{88, 2, 77, 73, 7, 63, 67},
	{99, 65, 4, 28, 6, 16, 70, 92},
	{41, 41, 26, 56, 83, 40, 80, 70, 33},
	{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
	{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
	{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
	{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
	{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
	{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23},
}

func solve() int {
	// Dynamic programming: work from bottom up
	var dp [rows]int
	for i := 0; i < rows; i++ {
		dp[i] = triangle[rows-1][i]
	}

	for row := rows - 2; row >= 0; row-- {
		for col := 0; col <= row; col++ {
			if dp[col] > dp[col+1] {
				dp[col] = triangle[row][col] + dp[col]
			} else {
				dp[col] = triangle[row][col] + dp[col+1]
			}
		}
	}

	return dp[0]
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

	fmt.Println("Problem 18: Maximum Path Sum I")
	fmt.Println("===============================")
	fmt.Printf("Finding maximum path sum in triangle, Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
