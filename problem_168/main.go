// Problem 168: Number Rotations
// Answer: 59206

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

const mod168 = 100000

func solve() int64 {
	total := int64(0)

	for k := 1; k <= 9; k++ {
		D := 10*k - 1
		for a0 := k; a0 <= 9; a0++ {
			// gcd(D, a0)
			g, tmpG := D, a0
			for tmpG != 0 {
				g, tmpG = tmpG, g%tmpG
			}
			Dp := D / g

			// Find order of 10 mod Dp
			ord := 0
			if Dp == 1 {
				ord = 1
			} else {
				pw := int64(1)
				for e := 1; e <= 10000; e++ {
					pw = pw * 10 % int64(Dp)
					if pw == 1 {
						ord = e
						break
					}
				}
				if ord == 0 {
					ord = Dp
				}
			}

			bigMod := int64(D) * mod168
			p10 := int64(10) % int64(Dp)
			p10Full := int64(10) % bigMod

			for d := 1; d <= 100; d++ {
				if d >= 2 {
					if Dp == 1 || p10 == 1 {
						pdMinus1 := (p10Full - 1 + bigMod) % bigMod
						num := (int64(a0) * pdMinus1) % bigMod
						nMod := (num / int64(D)) % mod168
						total = (total + nMod) % mod168
					}
				}
				p10 = p10 * 10 % int64(Dp)
				p10Full = p10Full * 10 % bigMod
			}
			_ = ord
		}
	}
	return total
}

func main() { bench.Run(168, solve) }
