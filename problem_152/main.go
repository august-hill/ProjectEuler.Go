// Problem 152: Writing 1/2 as a Sum of Inverse Squares
// Answer: 301

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

// All candidates have prime factors only from {2,3,5,7,13}.
// LCM of all k^2 for candidates = 2^12 * 3^6 * 5^4 * 7^2 * 13^2 = 15454353240000
// This fits in int64. We represent fractions as numerator / LCM.

var candList152 = []int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 15, 16, 18, 20, 21, 24, 27, 28,
	30, 32, 35, 36, 39, 40, 42, 45, 48, 52, 54, 56, 60, 63, 64, 65, 70, 72, 80,
}

const lcm152 = 2985984 * 625 * 49 * 169 // 2^12*3^6 * 5^4 * 7^2 * 13^2 = 15454353240000

var (
	invSqNum152  [39]int64 // lcm / k^2 for each candidate
	suffixSum152 [40]int64 // prefix sum from end
	resultCount  int
	candCount152 int
)

func gcd152(a, b int64) int64 {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func dfs152(idx int, remaining int64) {
	if remaining == 0 {
		resultCount++
		return
	}
	if remaining < 0 {
		return
	}
	if idx >= candCount152 {
		return
	}
	if remaining > suffixSum152[idx] {
		return
	}
	newRem := remaining - invSqNum152[idx]
	if newRem >= 0 {
		dfs152(idx+1, newRem)
	}
	dfs152(idx+1, remaining)
}

func solve() int64 {
	candCount152 = len(candList152)
	for i, k := range candList152 {
		k2 := int64(k) * int64(k)
		invSqNum152[i] = lcm152 / k2
	}
	suffixSum152[candCount152] = 0
	for i := candCount152 - 1; i >= 0; i-- {
		suffixSum152[i] = suffixSum152[i+1] + invSqNum152[i]
	}
	resultCount = 0
	// target = 1/2 = lcm/2
	dfs152(0, lcm152/2)
	return int64(resultCount)
}

func main() { bench.Run(152, solve) }
