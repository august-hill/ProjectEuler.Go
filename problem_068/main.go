// Problem 068: Magic 5-gon Ring
// Find the maximum 16-digit string for a "magic" 5-gon ring.
// Answer: 6531031914842725

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

func nextPermutation(arr []int) bool {
	n := len(arr)
	i := n - 1
	for i > 0 && arr[i-1] >= arr[i] {
		i--
	}
	if i == 0 {
		return false
	}
	j := n - 1
	for arr[j] <= arr[i-1] {
		j--
	}
	arr[i-1], arr[j] = arr[j], arr[i-1]
	// reverse from i to end
	for l, r := i, n-1; l < r; l, r = l+1, r-1 {
		arr[l], arr[r] = arr[r], arr[l]
	}
	return true
}

func solve() int64 {
	perm := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	best := int64(0)

	for {
		outer := perm[:5]
		inner := perm[5:]

		target := outer[0] + inner[0] + inner[1]
		valid := true
		for i := 1; i < 5; i++ {
			if outer[i]+inner[i]+inner[(i+1)%5] != target {
				valid = false
				break
			}
		}

		if valid {
			// Find line with smallest outer node
			minIdx := 0
			for i := 1; i < 5; i++ {
				if outer[i] < outer[minIdx] {
					minIdx = i
				}
			}

			// Build digit sequence
			digits := make([]int, 0, 15)
			for k := 0; k < 5; k++ {
				i := (minIdx + k) % 5
				digits = append(digits, outer[i], inner[i], inner[(i+1)%5])
			}

			// Count total digits
			totalDigits := 0
			for _, d := range digits {
				if d >= 10 {
					totalDigits += 2
				} else {
					totalDigits++
				}
			}

			if totalDigits == 16 {
				num := int64(0)
				for _, d := range digits {
					if d >= 10 {
						num = num*100 + int64(d)
					} else {
						num = num*10 + int64(d)
					}
				}
				if num > best {
					best = num
				}
			}
		}

		if !nextPermutation(perm) {
			break
		}
	}

	return best
}

func main() { bench.Run(68, solve) }
