// Problem 130: Composites with Prime Repunit Property
// Find the sum of the first 25 composite n with (n-1) divisible by A(n).
// Answer: 149253

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func repunitDiv130(n int) int {
	r, k := 1, 1
	for r%n != 0 {
		r = (r*10 + 1) % n
		k++
	}
	return k
}

func isPrime130(n int) bool {
	if n < 2 {
		return false
	}
	if n < 4 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func solve() int64 {
	var sum int64
	count := 0
	for n := 2; count < 25; n++ {
		if n%2 == 0 || n%5 == 0 {
			continue
		}
		if isPrime130(n) {
			continue
		}
		a := repunitDiv130(n)
		if (n-1)%a == 0 {
			sum += int64(n)
			count++
		}
	}
	return sum
}

func main() { bench.Run(130, solve) }
