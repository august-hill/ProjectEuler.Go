// Problem 77: Prime Summations
// What is the first value which can be written as the sum of primes in over five thousand different ways?
// Answer: 71

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

const sieveLimit077 = 1000

func solve() int64 {
	// Sieve of Eratosthenes
	isPrime := make([]bool, sieveLimit077)
	for i := 2; i < sieveLimit077; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i < sieveLimit077; i++ {
		if isPrime[i] {
			for j := i * i; j < sieveLimit077; j += i {
				isPrime[j] = false
			}
		}
	}

	primes := make([]int, 0, 200)
	for i := 2; i < sieveLimit077; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	for target := 2; target < sieveLimit077; target++ {
		dp := make([]int64, target+1)
		dp[0] = 1

		for _, p := range primes {
			if p > target {
				break
			}
			for i := p; i <= target; i++ {
				dp[i] += dp[i-p]
			}
		}

		if dp[target] > 5000 {
			return int64(target)
		}
	}

	return 0
}

func main() { bench.Run(77, solve) }
