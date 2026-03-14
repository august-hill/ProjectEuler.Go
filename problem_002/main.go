// Problem 002: Even Fibonacci Numbers
// Find the sum of all even-valued Fibonacci terms below 4 million.

package main

import (
	"fmt"
	"time"
)

const limit = 4_000_000

// sieve generates all Fibonacci numbers and filters for even ones
func sieve(n int) int {
	sum := 0
	a, b := 1, 2
	for b < n {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}
	return sum
}

// evenRecurrence uses E(n) = 4*E(n-1) + E(n-2) to generate only even Fibs
func evenRecurrence(n int) int {
	sum := 0
	a, b := 2, 8
	for a < n {
		sum += a
		a, b = b, 4*b+a
	}
	return sum
}

// closedForm uses the identity: Sum of F(3k) = (F(3n+2) - 1) / 2
func closedForm(n int) int {
	// First find the largest F(3k) < n
	// Generate Fibonacci sequence to find F(3n+2)
	a, b := 1, 1
	lastEvenIdx := 0
	idx := 2

	for {
		a, b = b, a+b
		idx++
		if idx%3 == 0 && b >= n {
			break
		}
		if idx%3 == 0 {
			lastEvenIdx = idx
		}
	}

	// Now compute F(lastEvenIdx + 2)
	a, b = 1, 1
	for i := 2; i < lastEvenIdx+2; i++ {
		a, b = b, a+b
	}

	return (b - 1) / 2
}

// matrixClosedForm uses matrix exponentiation for O(log n) Fibonacci
func matrixClosedForm(n int) int {
	// Find k where F(3k) < n but F(3k+3) >= n
	k := 1
	for fib(3*k) < n {
		k++
	}
	k--

	// Sum = (F(3k+2) - 1) / 2
	return (fib(3*k+2) - 1) / 2
}

// fib computes F(n) using matrix exponentiation in O(log n)
func fib(n int) int {
	if n <= 1 {
		return n
	}
	// Matrix [[1,1],[1,0]]^n gives [[F(n+1),F(n)],[F(n),F(n-1)]]
	a, b, c, d := 1, 1, 1, 0
	ra, rb, rc, rd := 1, 0, 0, 1 // Identity matrix

	for n > 0 {
		if n%2 == 1 {
			ra, rb, rc, rd = ra*a+rb*c, ra*b+rb*d, rc*a+rd*c, rc*b+rd*d
		}
		a, b, c, d = a*a+b*c, a*b+b*d, c*a+d*c, c*b+d*d
		n /= 2
	}
	return rb
}

func main() {
	fmt.Println("Problem 002: Even Fibonacci Numbers")
	fmt.Printf("Limit: %d\n\n", limit)

	start := time.Now()
	result1 := sieve(limit)
	elapsed1 := time.Since(start)

	start = time.Now()
	result2 := evenRecurrence(limit)
	elapsed2 := time.Since(start)

	start = time.Now()
	result3 := closedForm(limit)
	elapsed3 := time.Since(start)

	start = time.Now()
	result4 := matrixClosedForm(limit)
	elapsed4 := time.Since(start)

	fmt.Printf("Sieve (generate+filter): %d  (%v)\n", result1, elapsed1)
	fmt.Printf("Even Recurrence:         %d  (%v)\n", result2, elapsed2)
	fmt.Printf("Closed Form (iterative): %d  (%v)\n", result3, elapsed3)
	fmt.Printf("Closed Form (matrix):    %d  (%v)\n", result4, elapsed4)

	if result1 != result2 || result2 != result3 || result3 != result4 {
		fmt.Println("\nWARNING: Results do not match!")
	}
}
