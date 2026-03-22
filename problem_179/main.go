// Problem 179: Consecutive Positive Divisors
// Answer: 986262

package main

import (
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const limit179 = 10000000

var (
	once179     sync.Once
	divcount179 [limit179 + 1]int
)

func init179() {
	for i := 1; i <= limit179; i++ {
		for j := i; j <= limit179; j += i {
			divcount179[j]++
		}
	}
}

func solve() int64 {
	once179.Do(init179)
	count := int64(0)
	for n := 2; n < limit179; n++ {
		if divcount179[n] == divcount179[n+1] {
			count++
		}
	}
	return count
}

func main() { bench.Run(179, solve) }
