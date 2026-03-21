// Problem 088: Product-sum Numbers
// Find the sum of all minimal product-sum numbers for 2 <= k <= 12000.
// Answer: 7587457

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

const kMax088 = 12000
const psLimit088 = 2 * kMax088

var minPS088 [kMax088 + 1]int

func factorize088(product, sum, count, minFactor int) {
	k := product - sum + count
	if k <= kMax088 {
		if product < minPS088[k] {
			minPS088[k] = product
		}
	}

	for f := minFactor; f <= psLimit088; f++ {
		newProduct := product * f
		if newProduct > psLimit088 {
			break
		}
		factorize088(newProduct, sum+f, count+1, f)
	}
}

func solve() int64 {
	for i := 0; i <= kMax088; i++ {
		minPS088[i] = psLimit088 + 1
	}

	for f := 2; f <= psLimit088; f++ {
		factorize088(f, f, 1, f)
	}

	seen := make([]bool, psLimit088+2)
	total := 0
	for k := 2; k <= kMax088; k++ {
		v := minPS088[k]
		if !seen[v] {
			seen[v] = true
			total += v
		}
	}
	return int64(total)
}

func main() { bench.Run(88, solve) }
