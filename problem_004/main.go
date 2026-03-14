// Problem 004: Largest Palindrome Product
// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"fmt"
	"time"
)

// isPalindrome checks if a number is palindromic by reversing digits
func isPalindrome(n int) bool {
	if n < 0 {
		return false
	}
	original := n
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return original == reversed
}

// bruteForce iterates from high to low with early termination
func bruteForce() int {
	maxPalindrome := 0

	for i := 999; i >= 100; i-- {
		for j := i; j >= 100; j-- {
			product := i * j
			if product <= maxPalindrome {
				break // products only get smaller from here
			}
			if isPalindrome(product) {
				maxPalindrome = product
			}
		}
	}

	return maxPalindrome
}

// generatePalindromes generates 6-digit palindromes and checks if factorable
func generatePalindromes() int {
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
							return palindrome
						}
					}
				}
			}
		}
	}

	return 0
}

// divisibleBy11 uses the fact that all 6-digit palindromes are divisible by 11
func divisibleBy11() int {
	maxPalindrome := 0

	// One factor must be divisible by 11
	// So we iterate i over multiples of 11, j over all 3-digit numbers
	for i := 990; i >= 110; i -= 11 { // multiples of 11 from 990 down
		for j := 999; j >= i; j-- {
			product := i * j
			if product <= maxPalindrome {
				break
			}
			if isPalindrome(product) {
				maxPalindrome = product
			}
		}
	}

	return maxPalindrome
}

func main() {
	fmt.Println("Problem 004: Largest Palindrome Product")
	fmt.Println()

	start := time.Now()
	result1 := bruteForce()
	elapsed1 := time.Since(start)

	start = time.Now()
	result2 := generatePalindromes()
	elapsed2 := time.Since(start)

	start = time.Now()
	result3 := divisibleBy11()
	elapsed3 := time.Since(start)

	fmt.Printf("Brute Force (early term):   %d  (%v)\n", result1, elapsed1)
	fmt.Printf("Generate Palindromes:       %d  (%v)\n", result2, elapsed2)
	fmt.Printf("Divisible by 11:            %d  (%v)\n", result3, elapsed3)

	if result1 != result2 || result2 != result3 {
		fmt.Println("\nWARNING: Results do not match!")
	}
}
