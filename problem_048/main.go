// Problem 48: Self Powers
// Find last 10 digits of 1^1 + 2^2 + 3^3 + ... + 1000^1000.
// Answer: 9110846700

package main

import (
	"fmt"
	"math/big"
	"time"
)

func solve() int64 {
	mod := big.NewInt(10000000000) // 10^10
	sum := big.NewInt(0)

	for i := int64(1); i <= 1000; i++ {
		base := big.NewInt(i)
		power := new(big.Int).Exp(base, base, mod)
		sum.Add(sum, power)
		sum.Mod(sum, mod)
	}
	return sum.Int64()
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
	fmt.Println("Problem 48: Self Powers")
	fmt.Println("========================")
	benchmark(10000)
}
