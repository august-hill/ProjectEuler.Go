// Problem 125: Palindromic Sums
// Find sum of palindromes below 10^8 that are sums of consecutive squares.
// Answer: 2906969179

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const limit125 = 100000000

func isPalindrome125(n int64) bool {
	if n < 0 {
		return false
	}
	rev, orig := int64(0), n
	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	return rev == orig
}

func solve() int64 {
	const hashSize = 131072
	const hashMask = hashSize - 1
	hashTable := make([]int64, hashSize)
	hashUsed := make([]bool, hashSize)

	var total int64

	for i := int64(1); i*i < limit125; i++ {
		sum := i * i
		for j := i + 1; sum+j*j < limit125; j++ {
			sum += j * j
			if isPalindrome125(sum) {
				h := int(sum & hashMask)
				for hashUsed[h] {
					if hashTable[h] == sum {
						goto skip
					}
					h = (h + 1) & hashMask
				}
				hashTable[h] = sum
				hashUsed[h] = true
				total += sum
			skip:
			}
		}
	}

	return total
}

func main() { bench.Run(125, solve) }
