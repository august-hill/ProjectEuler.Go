// Problem 56: Powerful Digit Sum
// Considering natural numbers of the form a^b, where a, b < 100, find the maximum digital sum.
// Answer: 972

package main

import (
	"fmt"
	"math/big"
	"time"
)

func digitSum(n *big.Int) int {
	sum := 0
	for _, ch := range n.String() {
		sum += int(ch - '0')
	}
	return sum
}

func solve() int {
	maxSum := 0
	for a := 2; a < 100; a++ {
		base := big.NewInt(int64(a))
		power := big.NewInt(1)
		for b := 1; b < 100; b++ {
			power.Mul(power, base)
			s := digitSum(power)
			if s > maxSum {
				maxSum = s
			}
		}
	}
	return maxSum
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
	fmt.Println("Problem 56: Powerful Digit Sum")
	fmt.Println("==============================")
	benchmark(100)
}
