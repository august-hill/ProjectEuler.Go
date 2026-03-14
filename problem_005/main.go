// Problem 005: Smallest Multiple
// Find the smallest positive number divisible by all numbers from 1 to 20.

package main

import (
	"fmt"
	"time"
)

const limit = 20

// bruteForce checks multiples until one is divisible by all 1..n
func bruteForce(n int) int {
	// Start with n and increment by n (must be divisible by n)
	candidate := n
	for {
		divisible := true
		for i := n - 1; i >= 2; i-- {
			if candidate%i != 0 {
				divisible = false
				break
			}
		}
		if divisible {
			return candidate
		}
		candidate += n
	}
}

// gcd computes greatest common divisor using Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// lcm computes least common multiple
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

// iterativeLCM computes LCM of 1..n iteratively
func iterativeLCM(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result = lcm(result, i)
	}
	return result
}

// binaryGCD computes GCD using Stein's algorithm (binary GCD)
// Uses only subtraction and bit shifts, no division
func binaryGCD(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	// Find common factors of 2
	shift := 0
	for ((a | b) & 1) == 0 {
		a >>= 1
		b >>= 1
		shift++
	}

	// Remove remaining factors of 2 from a
	for (a & 1) == 0 {
		a >>= 1
	}

	for b != 0 {
		// Remove factors of 2 from b
		for (b & 1) == 0 {
			b >>= 1
		}

		// Now both a and b are odd
		if a > b {
			a, b = b, a
		}
		b = b - a
	}

	return a << shift
}

// binaryLCM computes LCM using binary GCD
func binaryLCM(a, b int) int {
	return a / binaryGCD(a, b) * b
}

// iterativeBinaryLCM computes LCM of 1..n using binary GCD
func iterativeBinaryLCM(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result = binaryLCM(result, i)
	}
	return result
}

func main() {
	fmt.Println("Problem 005: Smallest Multiple")
	fmt.Printf("Limit: %d\n\n", limit)

	start := time.Now()
	result1 := bruteForce(limit)
	elapsed1 := time.Since(start)

	start = time.Now()
	result2 := iterativeLCM(limit)
	elapsed2 := time.Since(start)

	start = time.Now()
	result3 := iterativeBinaryLCM(limit)
	elapsed3 := time.Since(start)

	fmt.Printf("Brute Force:          %d  (%v)\n", result1, elapsed1)
	fmt.Printf("Iterative LCM (GCD):  %d  (%v)\n", result2, elapsed2)
	fmt.Printf("Binary GCD variant:   %d  (%v)\n", result3, elapsed3)

	if result1 != result2 || result2 != result3 {
		fmt.Println("\nWARNING: Results do not match!")
	}
}
