// Problem 160: Factorial Trailing Digits
// Answer: 16576

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func powerMod160(base, exp, mod int64) int64 {
	result := int64(1)
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			result = result * base % mod
		}
		base = base * base % mod
		exp >>= 1
	}
	return result
}

func factmod160(n, p, pk int64) int64 {
	if n <= 1 {
		return 1
	}
	fullPeriodProd := int64(1)
	for i := int64(1); i <= pk; i++ {
		if i%p != 0 {
			fullPeriodProd = fullPeriodProd * (i % pk) % pk
		}
	}
	numFullPeriods := n / pk
	result := powerMod160(fullPeriodProd, numFullPeriods, pk)

	remainder := n % pk
	for i := int64(1); i <= remainder; i++ {
		if i%p != 0 {
			result = result * i % pk
		}
	}
	result = result * factmod160(n/p, p, pk) % pk
	return result
}

func countFactors160(n, p int64) int64 {
	count := int64(0)
	pk := p
	for pk <= n {
		count += n / pk
		pk *= p
	}
	return count
}

func solve() int64 {
	N := int64(1000000000000) // 10^12
	_ = countFactors160(N, 2) // v2 not directly needed (extra 2s cancel mod 32)
	v5 := countFactors160(N, 5)

	// Mod 5^5 = 3125
	r5 := factmod160(N, 5, 3125)
	// phi(3125) = 2500
	inv2mod3125 := powerMod160(2, 2499, 3125)
	fMod3125 := r5 * powerMod160(inv2mod3125, v5%2500, 3125) % 3125

	// Mod 2^5 = 32: v2-v5 >> 5 so 2^(v2-v5) mod 32 = 0
	// f mod 32 = 0

	// CRT: x ≡ 0 (mod 32), x ≡ fMod3125 (mod 3125)
	// x = 32*k where 32*k ≡ fMod3125 (mod 3125)
	inv32 := powerMod160(32, 2499, 3125)
	k := fMod3125 * inv32 % 3125
	result := 32 * k
	return result % 100000
}

func main() { bench.Run(160, solve) }
