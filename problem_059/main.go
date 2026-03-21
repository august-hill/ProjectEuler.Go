// Problem 59: XOR Decryption
// Decrypt the message encrypted with a 3-letter lowercase key using XOR,
// and find the sum of the ASCII values in the original text.
// Answer: 129448

package main

import (
	_ "embed"
	"strconv"
	"strings"
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

//go:embed p059_cipher.txt
var cipherData string

var cipher []int
var loadOnce sync.Once

func loadCipher() {
	loadOnce.Do(func() {
		parts := strings.Split(strings.TrimSpace(cipherData), ",")
		cipher = make([]int, len(parts))
		for i, p := range parts {
			n, _ := strconv.Atoi(strings.TrimSpace(p))
			cipher[i] = n
		}
	})
}

func solve() int64 {
	loadCipher()
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
	return int64(bestSum)
}

func main() { bench.Run(59, solve) }
