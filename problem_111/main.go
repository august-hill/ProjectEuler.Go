// Problem 111: Primes with Runs
// For 10-digit primes, find M(10,d) for each digit d then sum qualifying primes.
// Answer: 612407567715

package main

import (
	"math/bits"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func modPow111(base, exp, mod uint64) uint64 {
	result := uint64(1)
	base %= mod
	for exp > 0 {
		if exp&1 != 0 {
			hi, lo := bits.Mul64(result, base)
			result = hi%(mod) | 0 // need proper 128-bit mod
			_ = lo
			// Use manual 128-bit mod
			result = mod128Rem111(result, base, mod, hi, lo)
		}
		hi, lo := bits.Mul64(base, base)
		base = mod128Rem111(0, base, mod, hi, lo)
		exp >>= 1
	}
	return result
}

func mod128Rem111(rHi, rLo, mod, hi, lo uint64) uint64 {
	// compute hi:lo mod mod using bits.Div64
	if hi == 0 {
		return lo % mod
	}
	_, rem := bits.Div64(hi%mod, lo, mod)
	return rem
}

func modPow111v2(base, exp, mod uint64) uint64 {
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

func isPrime111(n uint64) bool {
	if n < 2 {
		return false
	}
	if n < 4 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	witnesses := []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	d := n - 1
	r := 0
	for d%2 == 0 {
		d /= 2
		r++
	}
	for _, a := range witnesses {
		if a >= n {
			continue
		}
		x := modPow111v2(a, d, n)
		if x == 1 || x == n-1 {
			continue
		}
		composite := true
		for i := 0; i < r-1; i++ {
			hi, lo := bits.Mul64(x, x)
			if hi == 0 {
				x = lo % n
			} else {
				_, x = bits.Div64(hi%n, lo, n)
			}
			if x == n-1 {
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

func sumPrimesWithRepeated111(d, nfree int, found *int) int64 {
	n := 10
	var total int64

	for mask := 0; mask < (1 << n); mask++ {
		if bits.OnesCount(uint(mask)) != nfree {
			continue
		}

		var freePos [10]int
		nf := 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				freePos[nf] = i
				nf++
			}
		}

		var digits [10]int
		for i := 0; i < n; i++ {
			digits[i] = d
		}

		assignment := make([]int, nf)
		for {
			valid := true
			for i := 0; i < nf; i++ {
				if assignment[i] == d {
					valid = false
					break
				}
			}

			if valid {
				for i := 0; i < n; i++ {
					if mask&(1<<i) == 0 {
						digits[i] = d
					}
				}
				for i := 0; i < nf; i++ {
					digits[freePos[i]] = assignment[i]
				}

				if digits[0] != 0 {
					var num uint64
					for i := 0; i < n; i++ {
						num = num*10 + uint64(digits[i])
					}
					if isPrime111(num) {
						total += int64(num)
						(*found)++
					}
				}
			}

			// Increment assignment
			carry := true
			for i := nf - 1; i >= 0 && carry; i-- {
				assignment[i]++
				if assignment[i] >= 10 {
					assignment[i] = 0
				} else {
					carry = false
				}
			}
			if carry {
				break
			}
		}
	}
	return total
}

func solve() int64 {
	var total int64
	for d := 0; d <= 9; d++ {
		for nfree := 0; nfree <= 9; nfree++ {
			found := 0
			s := sumPrimesWithRepeated111(d, nfree, &found)
			if found > 0 {
				total += s
				break
			}
		}
	}
	return total
}

func main() { bench.Run(111, solve) }
