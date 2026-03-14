// Problem 8: Largest Product in a Series
// Find the thirteen adjacent digits with the greatest product.

package main

import (
	"fmt"
	"time"
)

const digits = "73167176531330624919225119674426574742355349194934" +
	"96983520312774506326239578318016984801869478851843" +
	"85861560789112949495459501737958331952853208805511" +
	"12540698747158523863050715693290963295227443043557" +
	"66896648950445244523161731856403098711121722383113" +
	"62229893423380308135336276614282806444486645238749" +
	"30358907296290491560440772390713810515859307960866" +
	"70172427121883998797908792274921901699720888093776" +
	"65727333001053367881220235421809751254540594752243" +
	"52584907711670556013604839586446706324415722155397" +
	"53697817977846174064955149290862569321978468622482" +
	"83972241375657056057490261407972968652414535100474" +
	"82166370484403199890008895243450658541227588666881" +
	"16427171479924442928230863465674813919123162824586" +
	"17866458359124566529476545682848912883142607690042" +
	"24219022671055626321111109370544217506941658960408" +
	"07198403850962455444362981230987879927244284909188" +
	"84580156166097919133875499200524063689912560717606" +
	"05886116467109405077541002256983155200055935729725" +
	"71636269561882670428252483600823257530420752963450"

// Naive: O(n * k) - recalculate product for each window
func naive(s string, k int) uint64 {
	var maxProduct uint64 = 0
	n := len(s)

	for i := 0; i <= n-k; i++ {
		var product uint64 = 1
		for j := 0; j < k; j++ {
			product *= uint64(s[i+j] - '0')
		}
		if product > maxProduct {
			maxProduct = product
		}
	}
	return maxProduct
}

// Optimized: O(n) - split by zeros, use divide/multiply sliding
func optimized(s string, k int) uint64 {
	var maxProduct uint64 = 0
	n := len(s)
	start := 0

	for start <= n-k {
		// Find next zero-free segment
		zeroPos := -1
		for i := start; i < n && i < start+k; i++ {
			if s[i] == '0' {
				zeroPos = i
			}
		}

		if zeroPos != -1 {
			// Zero in initial window, skip past it
			start = zeroPos + 1
			continue
		}

		// Calculate initial product for this segment
		var product uint64 = 1
		for i := start; i < start+k; i++ {
			product *= uint64(s[i] - '0')
		}
		if product > maxProduct {
			maxProduct = product
		}

		// Slide through zero-free portion
		for i := start + k; i < n; i++ {
			if s[i] == '0' {
				// Hit a zero, restart after it
				start = i + 1
				break
			}
			// Divide by outgoing, multiply by incoming
			product = product / uint64(s[i-k]-'0') * uint64(s[i]-'0')
			if product > maxProduct {
				maxProduct = product
			}
			if i == n-1 {
				start = n // Done
			}
		}
	}
	return maxProduct
}

func benchmark(name string, f func(string, int) uint64, s string, k int, iterations int) time.Duration {
	// Warmup
	for i := 0; i < 100; i++ {
		f(s, k)
	}

	start := time.Now()
	for i := 0; i < iterations; i++ {
		f(s, k)
	}
	elapsed := time.Since(start)
	result := f(s, k)
	fmt.Printf("%s: %d (%.2f ns/op)\n", name, result, float64(elapsed.Nanoseconds())/float64(iterations))
	return elapsed
}

func main() {
	const k = 13
	const iterations = 100000

	fmt.Println("Problem 8: Largest Product in a Series")
	fmt.Println("=======================================")
	fmt.Printf("Window size: %d, Iterations: %d\n\n", k, iterations)

	naiveTime := benchmark("Naive    ", naive, digits, k, iterations)
	optTime := benchmark("Optimized", optimized, digits, k, iterations)

	fmt.Printf("\nSpeedup: %.2fx\n", float64(naiveTime)/float64(optTime))
}
