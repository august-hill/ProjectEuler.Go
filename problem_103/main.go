// Problem 103: Special Subset Sums: Optimum
// Find the optimum special subset sum set of size 7.
// Answer: 20313839404245

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func isSpecial103(set []int) bool {
	n := len(set)
	limit := 1 << n
	sums := make([]int, limit)
	sizes := make([]int, limit)

	for mask := 1; mask < limit; mask++ {
		s, sz := 0, 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				s += set[i]
				sz++
			}
		}
		sums[mask] = s
		sizes[mask] = sz
	}

	for a := 1; a < limit; a++ {
		for b := a + 1; b < limit; b++ {
			if a&b != 0 {
				continue
			}
			if sums[a] == sums[b] {
				return false
			}
			if sizes[a] > sizes[b] && sums[a] <= sums[b] {
				return false
			}
			if sizes[b] > sizes[a] && sums[b] <= sums[a] {
				return false
			}
		}
	}
	return true
}

func solve() int64 {
	base := []int{20, 31, 38, 39, 40, 42, 45}
	n := 7
	bestSum := 0
	for _, v := range base {
		bestSum += v
	}
	best := make([]int, n)
	copy(best, base)

	set := make([]int, n)
	for d0 := -3; d0 <= 3; d0++ {
		for d1 := -3; d1 <= 3; d1++ {
			for d2 := -3; d2 <= 3; d2++ {
				for d3 := -3; d3 <= 3; d3++ {
					for d4 := -3; d4 <= 3; d4++ {
						for d5 := -3; d5 <= 3; d5++ {
							for d6 := -3; d6 <= 3; d6++ {
								set[0] = base[0] + d0
								set[1] = base[1] + d1
								set[2] = base[2] + d2
								set[3] = base[3] + d3
								set[4] = base[4] + d4
								set[5] = base[5] + d5
								set[6] = base[6] + d6

								valid := true
								for i := 0; i < n && valid; i++ {
									if set[i] <= 0 {
										valid = false
									}
									for j := i + 1; j < n && valid; j++ {
										if set[i] >= set[j] {
											valid = false
										}
									}
								}
								if !valid {
									continue
								}

								s := 0
								for _, v := range set {
									s += v
								}
								if s >= bestSum {
									continue
								}

								if isSpecial103(set) {
									bestSum = s
									copy(best, set)
								}
							}
						}
					}
				}
			}
		}
	}

	var ans int64
	for _, v := range best {
		ans = ans*100 + int64(v)
	}
	return ans
}

func main() { bench.Run(103, solve) }
