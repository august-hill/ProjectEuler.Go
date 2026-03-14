// Problem 59: XOR Decryption
// Decrypt the message encrypted with a 3-letter lowercase key using XOR,
// and find the sum of the ASCII values in the original text.
// Answer: 129448

package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed p059_cipher.txt
var cipherData string

var cipher []int

func init() {
	parts := strings.Split(strings.TrimSpace(cipherData), ",")
	cipher = make([]int, len(parts))
	for i, p := range parts {
		n, _ := strconv.Atoi(strings.TrimSpace(p))
		cipher[i] = n
	}
}

func solve() int {
	bestSum := 0
	bestSpaces := 0

	for a := 'a'; a <= 'z'; a++ {
		for b := 'a'; b <= 'z'; b++ {
			for c := 'a'; c <= 'z'; c++ {
				key := [3]int{int(a), int(b), int(c)}
				sum := 0
				valid := true
				spaceCount := 0

				for i, val := range cipher {
					dec := val ^ key[i%3]
					if dec < 32 || dec > 126 {
						valid = false
						break
					}
					if dec == ' ' {
						spaceCount++
					}
					sum += dec
				}

				if valid && spaceCount > bestSpaces {
					bestSpaces = spaceCount
					bestSum = sum
				}
			}
		}
	}
	return bestSum
}

func benchmark(iterations int) time.Duration {
	for i := 0; i < 10; i++ {
		solve()
	}
	start := time.Now()
	var result int
	for i := 0; i < iterations; i++ {
		result = solve()
	}
	elapsed := time.Since(start)
	fmt.Printf("Result: %d (%.2f ns/op)\n", result, float64(elapsed.Nanoseconds())/float64(iterations))
	return elapsed
}

func main() {
	fmt.Println("Problem 59: XOR Decryption")
	fmt.Println("==========================")
	benchmark(100)
}
