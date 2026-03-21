// Problem 087: Prime Power Triples
// How many numbers below 50 million can be expressed as p^2 + q^3 + r^4?
// Answer: 1097343

package main

import (
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const limit087 = 50000000
const primeLimit087 = 7072

var (
	primes087   []int
	initOnce087 sync.Once
)

func initPrimes087() {
	initOnce087.Do(func() {
		sieve := make([]bool, primeLimit087+1)
		for i := 2; i <= primeLimit087; i++ {
			if !sieve[i] {
				primes087 = append(primes087, i)
				for j := i * i; j <= primeLimit087; j += i {
					sieve[j] = true
				}
			}
		}
	})
}

func solve() int64 {
	initPrimes087()

	seen := make([]bool, limit087)
	count := 0

	for _, r := range primes087 {
		r4 := r * r * r * r
		if r4 >= limit087 {
			break
		}
		for _, q := range primes087 {
			q3 := q * q * q
			if r4+q3 >= limit087 {
				break
			}
			for _, p := range primes087 {
				p2 := p * p
				total := r4 + q3 + p2
				if total >= limit087 {
					break
				}
				if !seen[total] {
					seen[total] = true
					count++
				}
			}
		}
	}

	return int64(count)
}

func main() { bench.Run(87, solve) }
