// Problem 118: Pandigital Prime Sets
// How many distinct sets of primes use all 9 digits exactly once?
// Answer: 44680

package main

import (
	"math/bits"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func modPow118(base, exp, mod uint64) uint64 {
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

func isPrime118(n uint64) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
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
		x := modPow118(a, d, n)
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

var totalCount118 int

func nextPerm118(arr []int) bool {
	n := len(arr)
	i := n - 2
	for i >= 0 && arr[i] >= arr[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := n - 1
	for arr[j] <= arr[i] {
		j--
	}
	arr[i], arr[j] = arr[j], arr[i]
	for a, b := i+1, n-1; a < b; a, b = a+1, b-1 {
		arr[a], arr[b] = arr[b], arr[a]
	}
	return true
}

func search118(usedMask int, prevNum int64) {
	if usedMask == 0x1FF {
		totalCount118++
		return
	}

	remaining := (^usedMask) & 0x1FF

	for sub := remaining; sub > 0; sub = (sub - 1) & remaining {
		var digs [9]int
		nd := 0
		for i := 0; i < 9; i++ {
			if sub&(1<<i) != 0 {
				digs[nd] = i + 1
				nd++
			}
		}

		perm := make([]int, nd)
		copy(perm, digs[:nd])

		for {
			var num int64
			for i := 0; i < nd; i++ {
				num = num*10 + int64(perm[i])
			}
			if num > prevNum && isPrime118(uint64(num)) {
				search118(usedMask|sub, num)
			}
			if !nextPerm118(perm) {
				break
			}
		}
	}
}

func solve() int64 {
	totalCount118 = 0
	search118(0, 0)
	return int64(totalCount118)
}

func main() { bench.Run(118, solve) }
