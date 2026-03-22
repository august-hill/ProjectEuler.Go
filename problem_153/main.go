// Problem 153: Investigating Gaussian Integers
// Answer: 17971254122360635

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

const n153 = 100000000

func sDiv153(M int64) int64 {
	if M <= 0 {
		return 0
	}
	result := int64(0)
	k := int64(1)
	for k <= M {
		q := M / k
		kMax := M / q
		sumK := (kMax-k+1) * (k + kMax) / 2
		result += q * sumK
		k = kMax + 1
	}
	return result
}

func gcd153(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func solve() int64 {
	realSum := sDiv153(n153)
	gaussSum := int64(0)
	for a := int64(1); a*a+1 <= n153; a++ {
		for b := int64(1); a*a+b*b <= n153; b++ {
			if gcd153(a, b) != 1 {
				continue
			}
			p := a*a + b*b
			M := int64(n153) / p
			gaussSum += 2 * a * sDiv153(M)
		}
	}
	return realSum + gaussSum
}

func main() { bench.Run(153, solve) }
