// Answer: 906609
// Problem 004: Largest Palindrome Product
// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

// generatePalindromes generates 6-digit palindromes and checks if factorable
func solve() int64 {
	// Generate palindromes from largest (999999) to smallest (100001)
	// 6-digit palindrome: abccba = 100001a + 10010b + 1100c
	for a := 9; a >= 1; a-- {
		for b := 9; b >= 0; b-- {
			for c := 9; c >= 0; c-- {
				palindrome := 100001*a + 10010*b + 1100*c

				// Check if it's a product of two 3-digit numbers
				// Only need to check up to sqrt(palindrome)
				for i := 999; i >= 100; i-- {
					if i*i < palindrome {
						break // i is too small, no valid j exists
					}
					if palindrome%i == 0 {
						j := palindrome / i
						if j >= 100 && j <= 999 {
							return int64(palindrome)
						}
					}
				}
			}
		}
	}

	return 0
}

func main() { bench.Run(4, solve) }
