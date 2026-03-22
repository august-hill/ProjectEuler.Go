// Problem 157: Solving the Diophantine equation 1/a + 1/b = p/10^n
// Answer: 53490

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func numDivisors157(n int64) int {
	count := 0
	for d := int64(1); d*d <= n; d++ {
		if n%d == 0 {
			count++
			if d != n/d {
				count++
			}
		}
	}
	return count
}

func gcd157(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func solve() int64 {
	total := int64(0)

	for n := 1; n <= 9; n++ {
		tenN := int64(1)
		for i := 0; i < n; i++ {
			tenN *= 10
		}

		var divs [200]int
		ndivs := 0
		for d := int64(1); d*d <= tenN; d++ {
			if tenN%d == 0 {
				divs[ndivs] = int(d)
				ndivs++
				if d != tenN/d {
					divs[ndivs] = int(tenN / d)
					ndivs++
				}
			}
		}

		for i := 0; i < ndivs; i++ {
			for j := i; j < ndivs; j++ {
				x := int64(divs[i])
				y := int64(divs[j])
				if x*y > tenN {
					continue
				}
				if tenN%(x*y) != 0 {
					continue
				}
				if gcd157(x, y) != 1 {
					continue
				}
				m := tenN / (x * y)
				D := m * (x + y)
				nd := numDivisors157(D)
				total += int64(nd)
			}
		}
	}
	return total
}

func main() { bench.Run(157, solve) }
