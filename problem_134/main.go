// Problem 134: Prime Pair Connection
// Find sum of S(p1, p2) for all prime pairs 5 <= p1 <= 10^6.
// Answer: 18613426663617118

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const limit134 = 1100000

func solve() int64 {
	notPrime := make([]bool, limit134)
	notPrime[0] = true
	notPrime[1] = true
	for i := 2; i*i < limit134; i++ {
		if !notPrime[i] {
			for j := i * i; j < limit134; j += i {
				notPrime[j] = true
			}
		}
	}

	var total int64

	for p1 := 5; p1 <= 1000000; {
		if notPrime[p1] {
			p1 += 2
			continue
		}
		p2 := p1 + 2
		for notPrime[p2] {
			p2 += 2
		}

		pow10 := int64(1)
		tmp := p1
		for tmp > 0 {
			pow10 *= 10
			tmp /= 10
		}

		// Find k such that k * pow10 ≡ -p1 (mod p2)
		// inv = pow10^(p2-2) mod p2
		base := pow10 % int64(p2)
		exp := int64(p2 - 2)
		mod := int64(p2)
		inv := int64(1)
		b := base
		e := exp
		for e > 0 {
			if e&1 != 0 {
				inv = inv * b % mod
			}
			b = b * b % mod
			e >>= 1
		}

		neg := (int64(-p1)%int64(p2) + int64(p2)) % int64(p2)
		k := neg * inv % int64(p2)
		n := int64(p1) + k*pow10
		total += n

		p1 = p2
	}
	return total
}

func main() { bench.Run(134, solve) }
