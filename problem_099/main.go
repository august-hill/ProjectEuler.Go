// Problem 099: Largest Exponential
// Using base/exponent pairs, find which line has the greatest numerical value.
// Answer: 709

package main

import (
	_ "embed"
	"math"
	"strconv"
	"strings"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

//go:embed p099_base_exp.txt
var baseExpData099 string

func solve() int64 {
	bestLine := 0
	bestVal := 0.0
	lineNum := 0

	lines := strings.Split(strings.TrimSpace(baseExpData099), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		lineNum++
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		base, _ := strconv.ParseFloat(parts[0], 64)
		exp, _ := strconv.ParseFloat(parts[1], 64)
		val := exp * math.Log(base)
		if val > bestVal {
			bestVal = val
			bestLine = lineNum
		}
	}

	return int64(bestLine)
}

func main() { bench.Run(99, solve) }
