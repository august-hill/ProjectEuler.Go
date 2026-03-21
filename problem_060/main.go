// Problem 060: Prime Pair Sets
// Find the lowest sum for a set of five primes where any two concatenate to produce another prime.
// Answer: 26033

package main

import (
	"math"
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const sieveSize = 100_000_000
const limit = 10000

var sieveData []bool // true = composite
var sieveOnce sync.Once

func initSieve() {
	sieveOnce.Do(func() {
		sieveData = make([]bool, sieveSize)
		sieveData[0] = true
		sieveData[1] = true
		for i := 2; i*i < sieveSize; i++ {
			if !sieveData[i] {
				for j := i * i; j < sieveSize; j += i {
					sieveData[j] = true
				}
			}
		}
	})
}

func modPow(base, exp, mod uint64) uint64 {
	result := uint64(1)
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			result = mulMod(result, base, mod)
		}
		exp >>= 1
		base = mulMod(base, base, mod)
	}
	return result
}

func mulMod(a, b, mod uint64) uint64 {
	// Use float64 trick for modular multiplication to avoid overflow
	if a < 1<<32 && b < 1<<32 {
		return (a * b) % mod
	}
	// Peasant multiplication
	result := uint64(0)
	a %= mod
	for b > 0 {
		if b&1 == 1 {
			result = (result + a) % mod
		}
		a = (a * 2) % mod
		b >>= 1
	}
	return result
}

func millerRabin(n uint64) bool {
	if n < 2 {
		return false
	}
	if n < 4 {
		return true
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

	witnesses := []uint64{2, 3, 5, 7}
	for _, a := range witnesses {
		if a >= n {
			continue
		}
		x := modPow(a, d, n)
		if x == 1 || x == n-1 {
			continue
		}
		composite := true
		for i := 0; i < r-1; i++ {
			x = mulMod(x, x, n)
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

func checkPrime(n int) bool {
	if n < sieveSize {
		return !sieveData[n]
	}
	return millerRabin(uint64(n))
}

func concatNums(a, b int) int {
	mult := 1
	tmp := b
	for tmp > 0 {
		mult *= 10
		tmp /= 10
	}
	return a*mult + b
}

func isPair(a, b int) bool {
	return checkPrime(concatNums(a, b)) && checkPrime(concatNums(b, a))
}

func solve() int64 {
	initSieve()
	primes := make([]int, 0, 1300)
	for i := 2; i < limit; i++ {
		if !sieveData[i] {
			primes = append(primes, i)
		}
	}
	n := len(primes)

	best := int(math.MaxInt32)

	for ai := 0; ai < n; ai++ {
		a := primes[ai]
		if a*5 >= best {
			break
		}

		for bi := ai + 1; bi < n; bi++ {
			b := primes[bi]
			if (a+b)*5/2 >= best {
				break
			}
			if !isPair(a, b) {
				continue
			}

			for ci := bi + 1; ci < n; ci++ {
				c := primes[ci]
				if (a+b+c)*5/3 >= best {
					break
				}
				if !isPair(a, c) || !isPair(b, c) {
					continue
				}

				for di := ci + 1; di < n; di++ {
					d := primes[di]
					partial := a + b + c + d
					if partial >= best {
						break
					}
					if !isPair(a, d) || !isPair(b, d) || !isPair(c, d) {
						continue
					}

					for ei := di + 1; ei < n; ei++ {
						e := primes[ei]
						sum := partial + e
						if sum >= best {
							break
						}
						if !isPair(a, e) || !isPair(b, e) || !isPair(c, e) || !isPair(d, e) {
							continue
						}
						best = sum
					}
				}
			}
		}
	}

	return int64(best)
}

func main() { bench.Run(60, solve) }
