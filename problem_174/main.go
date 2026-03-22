// Problem 174: Counting the number of "hollow" square laminae that can form one, two, ...
// distinct arrangements. How many values of t (tiles) have 1 <= N(t) <= 10?
// Answer: 209566

package main

import (
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const limit174 = 1000001

var (
	once174   sync.Once
	nt174     [limit174]int
)

func init174() {
	for n := int64(3); ; n++ {
		mMax := n - 2
		mMin := int64(2)
		if n%2 != 0 {
			mMin = 1
		}
		if n*n-mMax*mMax >= limit174 {
			break
		}
		for m := mMax; m >= mMin; m -= 2 {
			t := n*n - m*m
			if t >= limit174 {
				continue
			}
			nt174[t]++
		}
	}
}

func solve() int64 {
	once174.Do(init174)
	count := int64(0)
	for t := 1; t < limit174; t++ {
		if nt174[t] >= 1 && nt174[t] <= 10 {
			count++
		}
	}
	return count
}

func main() { bench.Run(174, solve) }
