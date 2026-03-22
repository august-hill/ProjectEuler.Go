// Problem 162: Hexadecimal Numbers
// Answer: 3D58725572C62302 (hex) = decimal value returned

package main

import (
	"math/big"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	// f(k) = 15*16^(k-1) - 15^k - 2*14*15^(k-1) + 2*14^k + 13*14^(k-1) - 13^k
	total := new(big.Int)
	pow16 := big.NewInt(1)
	pow15 := big.NewInt(1)
	pow14 := big.NewInt(1)
	pow13 := big.NewInt(1)
	b16 := big.NewInt(16)
	b15 := big.NewInt(15)
	b14 := big.NewInt(14)
	b13 := big.NewInt(13)

	tmp := new(big.Int)

	for k := 1; k <= 16; k++ {
		pow16.Mul(pow16, b16)
		pow15.Mul(pow15, b15)
		pow14.Mul(pow14, b14)
		pow13.Mul(pow13, b13)

		// 15*16^(k-1)
		t16km1 := new(big.Int).Div(pow16, b16)
		// 14*15^(k-1) * 2
		t15km1 := new(big.Int).Div(pow15, b15)
		// 13*14^(k-1)
		t14km1 := new(big.Int).Div(pow14, b14)

		fk := new(big.Int)
		fk.Set(tmp.Mul(big.NewInt(15), t16km1))
		fk.Sub(fk, pow15)
		fk.Sub(fk, tmp.Mul(big.NewInt(28), t15km1)) // 2*14
		fk.Add(fk, tmp.Mul(big.NewInt(2), pow14))
		fk.Add(fk, tmp.Mul(big.NewInt(13), t14km1))
		fk.Sub(fk, pow13)

		total.Add(total, fk)
	}

	return total.Int64()
}

func main() { bench.Run(162, solve) }
