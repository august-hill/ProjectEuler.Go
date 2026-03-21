// Answer: 73682
// Problem 31: Coin Sums
// How many different ways can 2 pounds be made using any number of coins?

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func solve() int64 {
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

	return int64(ways[target])
}

func main() { bench.Run(31, solve) }
