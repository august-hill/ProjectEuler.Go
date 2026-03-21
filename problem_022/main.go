// Answer: 871198282
// Problem 22: Names Scores
// Total of all name scores in the file.

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func nameValue(name string) int {
	sum := 0
	for _, ch := range name {
		sum += int(ch-'A') + 1
	}
	return sum
}

var names []string

func init() {
	data, err := os.ReadFile("names.txt")
	if err != nil {
		fmt.Println("Error reading names.txt:", err)
		os.Exit(1)
	}

	raw := strings.ReplaceAll(string(data), "\"", "")
	names = strings.Split(raw, ",")
}

func solve() int64 {
	sorted := make([]string, len(names))
	copy(sorted, names)
	sort.Strings(sorted)

	total := int64(0)
	for i, name := range sorted {
		total += int64(i+1) * int64(nameValue(name))
	}
	return total
}

func main() { bench.Run(22, solve) }
