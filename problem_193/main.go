// Problem 193: Squarefree Numbers
// Answer: 684465067343069

package main

import (
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const sieveLimit193 = 33554432 // 2^25

var (
	once193   sync.Once
	mu193     []int8
	primes193 []int
)

func init193() {
	mu193 = make([]int8, sieveLimit193)
	for i := range mu193 {
		mu193[i] = 1
	}
	mu193[0] = 0

	// Linear sieve to find primes
	smallestPF := make([]int, sieveLimit193)
	primes193 = make([]int, 0, sieveLimit193/10)

	for i := 2; i < sieveLimit193; i++ {
		if smallestPF[i] == 0 {
			primes193 = append(primes193, i)
			smallestPF[i] = i
		}
		for _, p := range primes193 {
			if p > smallestPF[i] || int64(i)*int64(p) >= sieveLimit193 {
				break
			}
			smallestPF[i*p] = p
		}
	}

	// Compute Mobius function
	for i := range mu193 {
		mu193[i] = 1
	}
	for _, p := range primes193 {
		pp := int64(p)
		for j := pp; j < sieveLimit193; j += pp {
			mu193[j] = -mu193[j]
		}
		p2 := pp * pp
		for j := p2; j < sieveLimit193; j += p2 {
			mu193[j] = 0
		}
	}
}

func solve() int64 {
	once193.Do(init193)

	N := int64(1) << 50
	count := int64(0)

	for k := int64(1); k < sieveLimit193; k++ {
		if mu193[k] == 0 {
			continue
		}
		k2 := k * k
		if k2 > N {
			break
		}
		count += int64(mu193[k]) * (N / k2)
	}
	return count
}

func main() { bench.Run(193, solve) }
