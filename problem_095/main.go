// Problem 095: Amicable Chains
// Find the smallest member of the longest amicable chain
// with no element exceeding one million.
// Answer: 14316

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

const chainLimit095 = 1000001

var sumDiv095 []int

func initSieve095() {
	sumDiv095 = make([]int, chainLimit095)
	for i := 2; i < chainLimit095; i++ {
		sumDiv095[i] = 1
	}
	for i := 2; i < chainLimit095; i++ {
		for j := 2 * i; j < chainLimit095; j += i {
			sumDiv095[j] += i
		}
	}
}

func solve() int64 {
	if sumDiv095 == nil {
		initSieve095()
	}
	sumDiv := sumDiv095

	visited := make([]bool, chainLimit095)
	bestLen := 0
	bestMin := 0

	chain := make([]int, chainLimit095)
	inChain := make([]bool, chainLimit095)

	for start := 2; start < chainLimit095; start++ {
		if visited[start] {
			continue
		}

		chainLen := 0
		n := start

		// Clear inChain for this iteration
		for n > 0 && n < chainLimit095 && !inChain[n] {
			inChain[n] = true
			chain[chainLen] = n
			chainLen++
			n = sumDiv[n]
		}

		// Check for cycle
		if n > 0 && n < chainLimit095 && inChain[n] {
			cycleStart := 0
			for i := 0; i < chainLen; i++ {
				if chain[i] == n {
					cycleStart = i
					break
				}
			}
			cycleLen := chainLen - cycleStart

			if cycleLen > bestLen {
				bestLen = cycleLen
				bestMin = chain[cycleStart]
				for i := cycleStart + 1; i < chainLen; i++ {
					if chain[i] < bestMin {
						bestMin = chain[i]
					}
				}
			}
		}

		for i := 0; i < chainLen; i++ {
			if chain[i] < chainLimit095 {
				visited[chain[i]] = true
				inChain[chain[i]] = false
			}
		}
	}

	return int64(bestMin)
}

func main() { bench.Run(95, solve) }
