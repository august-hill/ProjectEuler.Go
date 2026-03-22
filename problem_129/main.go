// Problem 129: Repunit Divisibility
// Find the least n > 10^6 for which A(n) > 10^6.
// Answer: 1000023

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func repunitDiv129(n int) int {
	r, k := 1, 1
	for r%n != 0 {
		r = (r*10 + 1) % n
		k++
	}
	return k
}

func solve() int64 {
	for n := 1000001; ; n++ {
		if n%2 == 0 || n%5 == 0 {
			continue
		}
		if repunitDiv129(n) > 1000000 {
			return int64(n)
		}
	}
}

func main() { bench.Run(129, solve) }
