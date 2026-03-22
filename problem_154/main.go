// Problem 154: Exploring Pascal's Pyramid
// Answer: 479742450

package main

import (
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const n154 = 200000

var (
	once154 sync.Once
	f2_154  [n154 + 1]int64
	f5_154  [n154 + 1]int64
)

func init154() {
	f2_154[0] = 0
	f5_154[0] = 0
	for i := 1; i <= n154; i++ {
		v := 0
		x := i
		for x%2 == 0 {
			v++
			x /= 2
		}
		f2_154[i] = f2_154[i-1] + int64(v)
		v = 0
		x = i
		for x%5 == 0 {
			v++
			x /= 5
		}
		f5_154[i] = f5_154[i-1] + int64(v)
	}
}

func solve() int64 {
	once154.Do(init154)

	total2 := f2_154[n154]
	total5 := f5_154[n154]
	count := int64(0)

	for a := 0; a <= n154/3; a++ {
		ra2 := total2 - f2_154[a]
		ra5 := total5 - f5_154[a]
		for b := a; b <= (n154-a)/2; b++ {
			c := n154 - a - b
			rem2 := ra2 - f2_154[b] - f2_154[c]
			rem5 := ra5 - f5_154[b] - f5_154[c]
			if rem2 >= 12 && rem5 >= 12 {
				if a == b && b == c {
					count++
				} else if a == b || b == c {
					count += 3
				} else {
					count += 6
				}
			}
		}
	}
	return count
}

func main() { bench.Run(154, solve) }
