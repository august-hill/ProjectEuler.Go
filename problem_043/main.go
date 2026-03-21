// Problem 43: Sub-string Divisibility
// Find sum of 0-9 pandigitals with substring divisibility.
// Answer: 16695334890

package main

import (
	"sort"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

// Generate next lexicographic permutation in-place, returns false when done
func nextPermutation(a []int) bool {
	n := len(a)
	if n < 2 {
		return false
	}

	// Find largest k where a[k] < a[k+1]
	k := n - 2
	for k >= 0 && a[k] >= a[k+1] {
		k--
	}
	if k < 0 {
		return false
	}

	// Find largest l > k where a[k] < a[l]
	l := n - 1
	for a[k] >= a[l] {
		l--
	}

	// Swap
	a[k], a[l] = a[l], a[k]

	// Reverse a[k+1:]
	for i, j := k+1, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return true
}

func pow10(n int) int64 {
	result := int64(1)
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}

func solve() int64 {
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	divisors := []int{2, 3, 5, 7, 11, 13, 17}
	var sum int64

	for {
		valid := true
		for i, div := range divisors {
			substr := digits[i+1]*100 + digits[i+2]*10 + digits[i+3]
			if substr%div != 0 {
				valid = false
				break
			}
		}

		if valid {
			var n int64
			for i, d := range digits {
				n += int64(d) * pow10(9-i)
			}
			sum += n
		}

		if !nextPermutation(digits) {
			break
		}
	}

	// Reset for next iteration
	sort.Ints(digits)
	return sum
}

func main() { bench.Run(43, solve) }
