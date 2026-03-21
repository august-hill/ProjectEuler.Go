// Problem 79: Passcode Derivation
// Given login attempts, find the shortest possible secret passcode.
// Answer: 73162890

package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

var (
	attempts079     [][3]int
	loadOnce079     sync.Once
)

func loadData079() {
	loadOnce079.Do(func() {
		data, err := os.ReadFile("p079_keylog.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not open p079_keylog.txt\n")
			os.Exit(1)
		}

		lines := strings.Split(strings.TrimSpace(string(data)), "\n")
		attempts079 = make([][3]int, 0, len(lines))
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if len(line) >= 3 {
				attempts079 = append(attempts079, [3]int{
					int(line[0] - '0'),
					int(line[1] - '0'),
					int(line[2] - '0'),
				})
			}
		}
	})
}

func solve() int64 {
	loadData079()

	var digits [10]bool
	var after [10][10]bool // after[a][b] = a must come before b

	for _, att := range attempts079 {
		a, b, c := att[0], att[1], att[2]
		digits[a] = true
		digits[b] = true
		digits[c] = true
		after[a][b] = true
		after[a][c] = true
		after[b][c] = true
	}

	// Topological sort
	resultDigits := make([]int, 0, 10)
	var used [10]bool

	totalDigits := 0
	for i := 0; i < 10; i++ {
		if digits[i] {
			totalDigits++
		}
	}

	for len(resultDigits) < totalDigits {
		for d := 0; d < 10; d++ {
			if !digits[d] || used[d] {
				continue
			}
			hasPred := false
			for o := 0; o < 10; o++ {
				if o != d && digits[o] && !used[o] && after[o][d] {
					hasPred = true
					break
				}
			}
			if !hasPred {
				resultDigits = append(resultDigits, d)
				used[d] = true
				break
			}
		}
	}

	result := int64(0)
	for _, d := range resultDigits {
		result = result*10 + int64(d)
	}
	return result
}

func main() { bench.Run(79, solve) }
