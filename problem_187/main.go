// Problem 187: Semiprimes
// Answer: 17427258

package main

import (
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const limit187 = 100000000

var (
	once187   sync.Once
	primes187 []int
)

func init187() {
	sieve := make([]bool, limit187)
	sieve[0] = true
	sieve[1] = true
	for i := 2; i*i < limit187; i++ {
		if !sieve[i] {
			for j := i * i; j < limit187; j += i {
				sieve[j] = true
			}
		}
	}
	for i := 2; i < limit187; i++ {
		if !sieve[i] {
			primes187 = append(primes187, i)
		}
	}
}

func solve() int64 {
	once187.Do(init187)

	count := int64(0)
	for i := 0; i < len(primes187); i++ {
		p := int64(primes187[i])
		if p*p >= limit187 {
			break
		}
		maxQ := int64(limit187-1) / p
		lo, hi, ans := i, len(primes187)-1, i-1
		for lo <= hi {
			mid := (lo + hi) / 2
			if int64(primes187[mid]) <= maxQ {
				ans = mid
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		if ans >= i {
			count += int64(ans - i + 1)
		}
	}
	return count
}

func main() { bench.Run(187, solve) }
