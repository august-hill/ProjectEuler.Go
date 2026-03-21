// Problem 74: Digit Factorial Chains
// How many chains with a starting number below one million contain exactly sixty non-repeating terms?
// Answer: 402

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

const cacheSize074 = 2200000

var factorials074 [10]int

func initFactorials074() {
	factorials074[0] = 1
	for i := 1; i <= 9; i++ {
		factorials074[i] = factorials074[i-1] * i
	}
}

func digitFactorialSum(n int) int {
	sum := 0
	for n > 0 {
		sum += factorials074[n%10]
		n /= 10
	}
	return sum
}

func solve() int64 {
	initFactorials074()
	chainLen := make([]uint8, cacheSize074)

	// Known loops
	chainLen[1] = 1
	chainLen[2] = 1
	chainLen[145] = 1
	chainLen[169] = 3
	chainLen[363601] = 3
	chainLen[1454] = 3
	chainLen[871] = 2
	chainLen[45361] = 2
	chainLen[872] = 2
	chainLen[45362] = 2

	count := 0
	chain := make([]int, 64)

	for start := 1; start < 1000000; start++ {
		chainIdx := 0
		n := start

		for {
			if n < cacheSize074 && chainLen[n] > 0 {
				remaining := int(chainLen[n])
				total := chainIdx + remaining

				for i := 0; i < chainIdx; i++ {
					l := total - i
					if chain[i] < cacheSize074 && l <= 255 {
						chainLen[chain[i]] = uint8(l)
					}
				}

				if total == 60 {
					count++
				}
				break
			}

			// Check if n is already in chain
			found := false
			for i := 0; i < chainIdx; i++ {
				if chain[i] == n {
					found = true
					break
				}
			}
			if found {
				break
			}

			chain[chainIdx] = n
			chainIdx++
			n = digitFactorialSum(n)
		}
	}

	return int64(count)
}

func main() { bench.Run(74, solve) }
