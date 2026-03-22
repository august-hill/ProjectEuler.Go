// Problem 122: Efficient Exponentiation
// Find the sum of minimum multiplications to compute n-th powers for n=1..200.
// Answer: 1582

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const maxN122 = 200

var best122 [maxN122 + 1]int

func dfs122(chain []int, length, maxDepth int) {
	cur := chain[length-1]
	if cur > maxN122 {
		return
	}
	if best122[cur] > length-1 {
		best122[cur] = length - 1
	}
	if length-1 >= maxDepth {
		return
	}
	for i := length - 1; i >= 0; i-- {
		next := cur + chain[i]
		if next > maxN122 || next <= cur {
			continue
		}
		chain[length] = next
		dfs122(chain, length+1, maxDepth)
	}
}

func solve() int64 {
	for i := 0; i <= maxN122; i++ {
		best122[i] = 100
	}
	best122[1] = 0

	chain := make([]int, 20)
	chain[0] = 1

	for depth := 1; depth <= 12; depth++ {
		dfs122(chain, 1, depth)
	}

	var total int64
	for i := 1; i <= maxN122; i++ {
		total += int64(best122[i])
	}
	return total
}

func main() { bench.Run(122, solve) }
