// Problem 200: Find the 200th prime-proof sqube containing the contiguous sub-string "200"
// Answer: 229161792008

package main

import (
	"sort"
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const sieveLim200 = 1000000

var (
	once200   sync.Once
	primes200 []int
)

func initSieve200() {
	sieve := make([]bool, sieveLim200)
	sieve[0] = true
	sieve[1] = true
	for i := 2; i*i < sieveLim200; i++ {
		if !sieve[i] {
			for j := i * i; j < sieveLim200; j += i {
				sieve[j] = true
			}
		}
	}
	primes200 = make([]int, 0, 80000)
	for i := 2; i < sieveLim200; i++ {
		if !sieve[i] {
			primes200 = append(primes200, i)
		}
	}
}

func modPow200(base, exp, mod uint64) uint64 {
	result := uint64(1)
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			result = result * base % mod
		}
		exp >>= 1
		base = base * base % mod
	}
	return result
}

func millerRabin200(n uint64) bool {
	if n < 2 {
		return false
	}
	if n < 4 {
		return n >= 2
	}
	if n%2 == 0 {
		return false
	}
	d := n - 1
	r := 0
	for d%2 == 0 {
		d /= 2
		r++
	}
	witnesses := []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	for _, a := range witnesses {
		if a >= n {
			continue
		}
		x := modPow200(a, d, n)
		if x == 1 || x == n-1 {
			continue
		}
		composite := true
		for i := 0; i < r-1; i++ {
			x = x * x % n
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

func contains200_200(n uint64) bool {
	var buf [20]byte
	length := 0
	tmp := n
	for tmp > 0 {
		buf[length] = byte(tmp % 10)
		length++
		tmp /= 10
	}
	for i := length - 1; i >= 2; i-- {
		if buf[i] == 2 && buf[i-1] == 0 && buf[i-2] == 0 {
			return true
		}
	}
	return false
}

func isPrimeProof200(n uint64) bool {
	var digits [20]byte
	length := 0
	tmp := n
	for tmp > 0 {
		digits[length] = byte(tmp % 10)
		length++
		tmp /= 10
	}

	for pos := 0; pos < length; pos++ {
		orig := digits[pos]
		for d := byte(0); d <= 9; d++ {
			if d == orig {
				continue
			}
			if pos == length-1 && d == 0 {
				continue
			}
			digits[pos] = d
			val := uint64(0)
			for i := length - 1; i >= 0; i-- {
				val = val*10 + uint64(digits[i])
			}
			if millerRabin200(val) {
				digits[pos] = orig
				return false
			}
		}
		digits[pos] = orig
	}
	return true
}

var (
	answerCache200 int64
)

func compute200() {
	once200.Do(initSieve200)

	limit := uint64(300000000000) // 3*10^11

	squbes := make([]uint64, 0, 1000000)

	for qi, q := range primes200 {
		qv := uint64(q)
		q3 := qv * qv * qv
		if q3 > limit {
			break
		}
		for pi, p := range primes200 {
			if pi == qi {
				continue
			}
			pv := uint64(p)
			p2 := pv * pv
			sqube := p2 * q3
			if sqube > limit {
				break
			}
			if contains200_200(sqube) {
				squbes = append(squbes, sqube)
			}
		}
	}

	sort.Slice(squbes, func(i, j int) bool { return squbes[i] < squbes[j] })

	// Deduplicate
	unique := squbes[:0]
	for i, v := range squbes {
		if i == 0 || v != squbes[i-1] {
			unique = append(unique, v)
		}
	}

	count := 0
	for _, v := range unique {
		if isPrimeProof200(v) {
			count++
			if count == 200 {
				answerCache200 = int64(v)
				return
			}
		}
	}
}

func solve() int64 {
	once200.Do(func() {})
	compute200()
	return answerCache200
}

func main() { bench.Run(200, solve) }
