// Problem 159: Digital Root Sums of Factorisations
// Answer: 14489159

package main

import (
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const limit159 = 1000000

var (
	once159    sync.Once
	mdrsArr159 [limit159]int
)

func digitalRoot159(n int) int {
	if n == 0 {
		return 0
	}
	r := n % 9
	if r == 0 {
		return 9
	}
	return r
}

func init159() {
	for i := 2; i < limit159; i++ {
		mdrsArr159[i] = digitalRoot159(i)
	}
	for i := 2; i < limit159; i++ {
		for j := 2; i*j < limit159; j++ {
			prod := i * j
			val := mdrsArr159[i] + mdrsArr159[j]
			if val > mdrsArr159[prod] {
				mdrsArr159[prod] = val
			}
		}
	}
}

func solve() int64 {
	once159.Do(init159)
	sum := int64(0)
	for i := 2; i < limit159; i++ {
		sum += int64(mdrsArr159[i])
	}
	return sum
}

func main() { bench.Run(159, solve) }
