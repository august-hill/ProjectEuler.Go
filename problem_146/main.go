// Problem 146: Investigating a Prime Pattern
// Find the sum of all n < 150,000,000 where n^2+1, +3, +7, +9, +13, +27 are consecutive primes.
// Answer: 676333270

package main

import (
	"math/bits"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func millerRabin146(n int64) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 || n == 5 || n == 7 {
		return true
	}
	if n%2 == 0 || n%3 == 0 || n%5 == 0 {
		return false
	}

	d := n - 1
	r := 0
	for d%2 == 0 {
		d /= 2
		r++
	}

	witnesses := []int64{2, 3, 5, 7, 11, 13}
	for _, a := range witnesses {
		if a >= n {
			continue
		}

		// Compute a^d mod n
		x := modPow146(uint64(a), uint64(d), uint64(n))
		xv := int64(x)
		if xv == 1 || xv == n-1 {
			continue
		}

		composite := true
		for i := 0; i < r-1; i++ {
			hi, lo := bits.Mul64(uint64(xv), uint64(xv))
			if hi == 0 {
				xv = int64(lo % uint64(n))
			} else {
				_, rem := bits.Div64(hi%uint64(n), lo, uint64(n))
				xv = int64(rem)
			}
			if xv == n-1 {
				composite = false
				break
			}
		}
		if composite {
			return false
		}
	}
	return true
}

func modPow146(base, exp, mod uint64) uint64 {
	result := uint64(1)
	base %= mod
	for exp > 0 {
		if exp&1 != 0 {
			hi, lo := bits.Mul64(result, base)
			if hi == 0 {
				result = lo % mod
			} else {
				_, result = bits.Div64(hi%mod, lo, mod)
			}
		}
		hi, lo := bits.Mul64(base, base)
		if hi == 0 {
			base = lo % mod
		} else {
			_, base = bits.Div64(hi%mod, lo, mod)
		}
		exp >>= 1
	}
	return result
}

func solve() int64 {
	const limit = 150000000
	var sum int64

	for n := int64(10); n < limit; n += 10 {
		if n%3 == 0 {
			continue
		}

		n7 := int(n % 7)
		n2_7 := (n7 * n7) % 7
		if (n2_7+1)%7 == 0 || (n2_7+3)%7 == 0 ||
			(n2_7+7)%7 == 0 || (n2_7+9)%7 == 0 ||
			(n2_7+13)%7 == 0 || (n2_7+27)%7 == 0 {
			continue
		}

		n13 := int(n % 13)
		n2_13 := (n13 * n13) % 13
		if (n2_13+1)%13 == 0 || (n2_13+3)%13 == 0 ||
			(n2_13+7)%13 == 0 || (n2_13+9)%13 == 0 ||
			(n2_13+13)%13 == 0 || (n2_13+27)%13 == 0 {
			continue
		}

		n2 := n * n
		if !millerRabin146(n2+1) || !millerRabin146(n2+3) ||
			!millerRabin146(n2+7) || !millerRabin146(n2+9) ||
			!millerRabin146(n2+13) || !millerRabin146(n2+27) {
			continue
		}
		if millerRabin146(n2+5) || millerRabin146(n2+11) ||
			millerRabin146(n2+15) || millerRabin146(n2+17) ||
			millerRabin146(n2+19) || millerRabin146(n2+21) ||
			millerRabin146(n2+23) || millerRabin146(n2+25) {
			continue
		}

		sum += n
	}
	return sum
}

func main() { bench.Run(146, solve) }
