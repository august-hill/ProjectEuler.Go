// Problem 098: Anagramic Squares
// Find the largest square number formed by anagram word pairs.
// Answer: 18769

package main

import (
	_ "embed"
	"math"
	"strings"
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

//go:embed p098_words.txt
var wordsData098 string

var (
	words098    []string
	initOnce098 sync.Once
)

func loadWords098() {
	initOnce098.Do(func() {
		words098 = nil
		parts := strings.Split(wordsData098, ",")
		for _, p := range parts {
			p = strings.TrimSpace(p)
			p = strings.Trim(p, "\"")
			if p != "" {
				words098 = append(words098, p)
			}
		}
	})
}

func areAnagrams098(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var counts [26]int
	for i := 0; i < len(a); i++ {
		counts[a[i]-'A']++
	}
	for i := 0; i < len(b); i++ {
		counts[b[i]-'A']--
	}
	for _, c := range counts {
		if c != 0 {
			return false
		}
	}
	return true
}

func isqrtCheck098(n int64) int64 {
	r := int64(math.Sqrt(float64(n)))
	if r*r == n {
		return r
	}
	if (r+1)*(r+1) == n {
		return r + 1
	}
	if r > 0 && (r-1)*(r-1) == n {
		return r - 1
	}
	return -1
}

func solve() int64 {
	loadWords098()

	var best int64

	for i := 0; i < len(words098); i++ {
		for j := i + 1; j < len(words098); j++ {
			if !areAnagrams098(words098[i], words098[j]) {
				continue
			}

			wlen := len(words098[i])
			loSq := int64(1)
			for k := 1; k < wlen; k++ {
				loSq *= 10
			}
			hiSq := loSq*10 - 1

			lo := int64(math.Ceil(math.Sqrt(float64(loSq))))
			hi := int64(math.Floor(math.Sqrt(float64(hiSq))))

			for s := lo; s <= hi; s++ {
				sq := s * s

				// Map words098[i] -> sq
				var letterToDigit [26]int
				var digitToLetter [10]int
				for k := range letterToDigit {
					letterToDigit[k] = -1
				}
				for k := range digitToLetter {
					digitToLetter[k] = -1
				}

				valid := true
				tmp := sq
				digits := make([]int, wlen)
				for k := wlen - 1; k >= 0; k-- {
					digits[k] = int(tmp % 10)
					tmp /= 10
				}

				for k := 0; k < wlen; k++ {
					li := int(words098[i][k] - 'A')
					di := digits[k]
					if letterToDigit[li] == -1 && digitToLetter[di] == -1 {
						letterToDigit[li] = di
						digitToLetter[di] = li
					} else if letterToDigit[li] != di || digitToLetter[di] != li {
						valid = false
						break
					}
				}
				if !valid {
					continue
				}

				// Apply mapping to words098[j]
				var num2 int64
				for k := 0; k < wlen; k++ {
					li := int(words098[j][k] - 'A')
					if letterToDigit[li] == -1 {
						valid = false
						break
					}
					num2 = num2*10 + int64(letterToDigit[li])
				}
				if !valid {
					continue
				}

				// No leading zero
				if letterToDigit[int(words098[j][0]-'A')] == 0 {
					continue
				}

				// Check if num2 is perfect square
				if isqrtCheck098(num2) >= 0 {
					mx := sq
					if num2 > mx {
						mx = num2
					}
					if mx > best {
						best = mx
					}
				}
			}
		}
	}

	return best
}

func main() { bench.Run(98, solve) }
