// Problem 120: Square Remainders
// Find the sum of rmax for 3 <= a <= 1000.
// Answer: 333082500

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func solve() int64 {
	var total int64
	for a := 3; a <= 1000; a++ {
		a2 := int64(a) * int64(a)
		var maxR int64
		for n := 1; n < 2*a; n += 2 {
			r := (2 * int64(n) * int64(a)) % a2
			if r > maxR {
				maxR = r
			}
		}
		total += maxR
	}
	return total
}

func main() { bench.Run(120, solve) }
