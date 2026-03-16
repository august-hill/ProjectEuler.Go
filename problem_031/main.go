// Problem 31: Coin Sums
// How many different ways can 2 pounds be made using any number of coins?
// Answer: 73682

package main

import (
	"fmt"
	"time"
)

func solve() int {
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	target := 200

	// Dynamic programming
	ways := make([]int, target+1)
	ways[0] = 1

	for _, coin := range coins {
		for amount := coin; amount <= target; amount++ {
			ways[amount] += ways[amount-coin]
		}
	}

	return ways[target]
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

	fmt.Println("Problem 31: Coin Sums")
	fmt.Println("======================")
	fmt.Printf("Ways to make 2 pounds with coins, Iterations: %d\n\n", iterations)

	benchmark(iterations)
}
