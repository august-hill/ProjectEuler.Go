// Problem 128: Hexagonal Tile Differences
// Find the 2000th tile with PD(tile) = 3.
// Answer: 14516824220

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func isPrime128(n int64) bool {
	if n < 2 {
		return false
	}
	if n < 4 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := int64(5); i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func solve() int64 {
	count := 1 // tile 1 has PD=3

	for r := int64(1); ; r++ {
		// First tile of ring r
		if isPrime128(6*r-1) && isPrime128(6*r+1) && isPrime128(12*r+5) {
			count++
			if count == 2000 {
				return 3*r*r - 3*r + 2
			}
		}

		// Last tile of ring r (valid for r >= 2)
		if r >= 2 {
			if isPrime128(6*r-1) && isPrime128(6*r+5) && isPrime128(12*r-7) {
				count++
				if count == 2000 {
					return 3*r*r + 3*r + 1
				}
			}
		}
	}
}

func main() { bench.Run(128, solve) }
