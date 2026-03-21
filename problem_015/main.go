// Answer: 137846528640
// Problem 15: Lattice Paths
// Count routes through 20x20 grid (only right/down moves).

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

// C(2n, n) computed without overflow by interleaving multiply/divide
func latticePaths(n uint64) uint64 {
	result := uint64(1)
	for i := uint64(1); i <= n; i++ {
		result = result * (n + i) / i
	}
	return result
}

func solve() int64 {
	return int64(latticePaths(20))
}

func main() { bench.Run(15, solve) }
