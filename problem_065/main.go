// Problem 065: Convergents of e
// Find the sum of digits in the numerator of the 100th convergent of e.
// Answer: 272

package main

import (
	"math/big"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func cfCoeff(k int) int64 {
	if k == 0 {
		return 2
	}
	j := k - 1
	if j%3 == 1 {
		return int64(2 * (j/3 + 1))
	}
	return 1
}

func solve() int64 {
	hPrev2 := big.NewInt(1) // h_{-1}
	hPrev1 := big.NewInt(2) // h_0

	for k := 1; k < 100; k++ {
		a := cfCoeff(k)
		mul := new(big.Int).Mul(hPrev1, big.NewInt(a))
		newH := new(big.Int).Add(mul, hPrev2)
		hPrev2 = hPrev1
		hPrev1 = newH
	}

	sum := 0
	for _, c := range hPrev1.String() {
		sum += int(c - '0')
	}
	return int64(sum)
}

func main() { bench.Run(65, solve) }
