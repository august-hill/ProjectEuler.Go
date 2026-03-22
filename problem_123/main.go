// Problem 123: Prime Square Remainders
// Find the least prime index n where (p_n-1)^n + (p_n+1)^n mod p_n^2 > 10^10.
// Answer: 21035

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const sieveSize123 = 1000000

func solve() int64 {
	sieve := make([]bool, sieveSize123)
	for i := range sieve {
		sieve[i] = true
	}
	sieve[0] = false
	sieve[1] = false
	for i := 2; int64(i)*int64(i) < sieveSize123; i++ {
		if sieve[i] {
			for j := i * i; j < sieveSize123; j += i {
				sieve[j] = false
			}
		}
	}

	primes := make([]int64, 0, 100000)
	for i := 2; i < sieveSize123; i++ {
		if sieve[i] {
			primes = append(primes, int64(i))
		}
	}

	target := int64(10000000000)
	for i, p := range primes {
		n := int64(i + 1)
		if n%2 == 0 {
			continue
		}
		p2 := p * p
		r := (2 * n * p) % p2
		if r > target {
			return n
		}
	}
	return -1
}

func main() { bench.Run(123, solve) }
