// Problem 092: Square Digit Chains
// How many starting numbers below ten million will arrive at 89?
// Answer: 8581146

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func squareDigitSum092(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

func solve() int64 {
	// Cache for small numbers (max square digit sum for 7 digits: 7 * 81 = 567)
	var cache [568]byte // 0 = unknown, 1 = ends at 1, 89 = ends at 89

	count := 0

	for n := 1; n < 10000000; n++ {
		chain := n

		for {
			if chain == 1 {
				break
			}
			if chain == 89 {
				count++
				break
			}
			if chain < 568 && cache[chain] != 0 {
				if cache[chain] == 89 {
					count++
				}
				break
			}
			chain = squareDigitSum092(chain)
		}

		// Cache for small n
		if n < 568 {
			if chain == 1 || (chain < 568 && cache[chain] == 1) {
				cache[n] = 1
			} else {
				cache[n] = 89
			}
		}
	}

	return int64(count)
}

func main() { bench.Run(92, solve) }
