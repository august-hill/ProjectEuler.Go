// Problem 042: Coded Triangle Numbers
// How many words in the file are triangle words?
// Answer: 162

package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

var words []string
var loadOnce sync.Once

func loadWords() {
	loadOnce.Do(func() {
		data, err := os.ReadFile("words.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading words.txt: %v\n", err)
			os.Exit(1)
		}
		parts := strings.Split(string(data), ",")
		for _, p := range parts {
			w := strings.Trim(p, "\" \n\r")
			if w != "" {
				words = append(words, w)
			}
		}
	})
}

func isTriangle(n int) bool {
	// n = k*(k+1)/2 => k^2 + k - 2n = 0 => k = (-1 + sqrt(1+8n))/2
	// Check if 1+8n is a perfect square
	val := 1 + 8*n
	k := 1
	for k*k < val {
		k++
	}
	return k*k == val
}

func wordValue(word string) int {
	sum := 0
	for _, c := range word {
		sum += int(c-'A') + 1
	}
	return sum
}

func solve() int64 {
	loadWords()
	count := 0
	for _, w := range words {
		if isTriangle(wordValue(w)) {
			count++
		}
	}
	return int64(count)
}

func main() { bench.Run(42, solve) }
