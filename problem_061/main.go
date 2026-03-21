// Problem 061: Cyclical Figurate Numbers
// Find the sum of the only ordered set of six cyclic 4-digit figurate numbers.
// Answer: 28684

package main

import (
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const maxPerPrefix = 10

var byPrefix [6][100][]int
var figurateOnce sync.Once

func initFigurates() {
	figurateOnce.Do(func() {
		for s := 3; s <= 8; s++ {
			idx := s - 3
			for i := range byPrefix[idx] {
				byPrefix[idx][i] = nil
			}
			for n := 1; ; n++ {
				val := n * ((s-2)*n - (s - 4)) / 2
				if val >= 10000 {
					break
				}
				if val >= 1000 {
					prefix := val / 100
					if len(byPrefix[idx][prefix]) < maxPerPrefix {
						byPrefix[idx][prefix] = append(byPrefix[idx][prefix], val)
					}
				}
			}
		}
	})
}

var chain [6]int
var usedType [6]bool

func search(depth int) int {
	if depth == 6 {
		if chain[5]%100 == chain[0]/100 {
			sum := 0
			for i := 0; i < 6; i++ {
				sum += chain[i]
			}
			return sum
		}
		return 0
	}

	for t := 0; t < 6; t++ {
		if usedType[t] {
			continue
		}
		usedType[t] = true

		if depth == 0 {
			for prefix := 10; prefix < 100; prefix++ {
				for _, val := range byPrefix[t][prefix] {
					if val%100 < 10 {
						continue
					}
					chain[0] = val
					result := search(1)
					if result > 0 {
						usedType[t] = false
						return result
					}
				}
			}
		} else {
			needed := chain[depth-1] % 100
			if needed < 10 {
				usedType[t] = false
				continue
			}
			for _, val := range byPrefix[t][needed] {
				if val%100 < 10 && depth < 5 {
					continue
				}
				chain[depth] = val
				result := search(depth + 1)
				if result > 0 {
					usedType[t] = false
					return result
				}
			}
		}

		usedType[t] = false
	}
	return 0
}

func solve() int64 {
	initFigurates()
	for i := range usedType {
		usedType[i] = false
	}
	return int64(search(0))
}

func main() { bench.Run(61, solve) }
