// Problem 173: Using up to one million tiles how many hollow square laminae can be formed?
// Answer: 1572729

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	count := int64(0)
	limit := int64(1000000)

	for n := int64(3); ; n++ {
		mMax := n - 2
		mMin := int64(2)
		if n%2 != 0 {
			mMin = 1
		}
		if n*n-mMax*mMax > limit {
			break
		}

		mSqMin := n*n - limit
		actualMMin := mMin
		if mSqMin > 0 {
			actualMMin = int64(math.Sqrt(float64(mSqMin)))
			if actualMMin*actualMMin < mSqMin {
				actualMMin++
			}
		}
		if actualMMin < mMin {
			actualMMin = mMin
		}
		if actualMMin%2 != n%2 {
			actualMMin++
		}
		if actualMMin > mMax {
			continue
		}
		count += (mMax-actualMMin)/2 + 1
	}
	return count
}

func main() { bench.Run(173, solve) }
