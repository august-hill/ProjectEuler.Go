// Problem 089: Roman Numerals
// Find how many characters are saved by writing each Roman numeral in minimal form.
// Answer: 743

package main

import (
	_ "embed"
	"strings"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

//go:embed p089_roman.txt
var romanData089 string

func romanCharVal089(c byte) int {
	switch c {
	case 'M':
		return 1000
	case 'D':
		return 500
	case 'C':
		return 100
	case 'L':
		return 50
	case 'X':
		return 10
	case 'V':
		return 5
	case 'I':
		return 1
	}
	return 0
}

func romanToInt089(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		v := romanCharVal089(s[i])
		if i+1 < len(s) {
			next := romanCharVal089(s[i+1])
			if v < next {
				total += next - v
				i++
				continue
			}
		}
		total += v
	}
	return total
}

func minimalRomanLen089(n int) int {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	lens := []int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1}
	length := 0
	for i, v := range vals {
		for n >= v {
			length += lens[i]
			n -= v
		}
	}
	return length
}

func solve() int64 {
	saved := 0
	lines := strings.Split(strings.TrimSpace(romanData089), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		value := romanToInt089(line)
		minLen := minimalRomanLen089(value)
		saved += len(line) - minLen
	}
	return int64(saved)
}

func main() { bench.Run(89, solve) }
