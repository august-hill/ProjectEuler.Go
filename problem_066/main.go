// Problem 066: Diophantine Equation
// Find D <= 1000 for which the minimal solution of x^2 - Dy^2 = 1 is largest in x.
// Answer: 661

package main

import (
	"math"
	"math/big"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	bestX := new(big.Int)
	bestD := 0

	one := big.NewInt(1)

	for d := 2; d <= 1000; d++ {
		a0 := int(math.Sqrt(float64(d)))
		if a0*a0 == d {
			continue
		}

		m, dn, a := 0, 1, a0

		hPrev2 := big.NewInt(1)
		hPrev1 := big.NewInt(int64(a0))
		kPrev2 := big.NewInt(0)
		kPrev1 := big.NewInt(1)

		bigD := big.NewInt(int64(d))

		for {
			m = dn*a - m
			dn = (d - m*m) / dn
			a = (a0 + m) / dn

			bigA := big.NewInt(int64(a))

			newH := new(big.Int).Mul(bigA, hPrev1)
			newH.Add(newH, hPrev2)

			newK := new(big.Int).Mul(bigA, kPrev1)
			newK.Add(newK, kPrev2)

			// Check if h^2 - D*k^2 = 1
			h2 := new(big.Int).Mul(newH, newH)
			k2 := new(big.Int).Mul(newK, newK)
			dk2 := new(big.Int).Mul(bigD, k2)
			dk2p1 := new(big.Int).Add(dk2, one)

			if h2.Cmp(dk2p1) == 0 {
				if newH.Cmp(bestX) > 0 {
					bestX.Set(newH)
					bestD = d
				}
				break
			}

			hPrev2 = hPrev1
			hPrev1 = newH
			kPrev2 = kPrev1
			kPrev1 = newK
		}
	}

	return int64(bestD)
}

func main() { bench.Run(66, solve) }
