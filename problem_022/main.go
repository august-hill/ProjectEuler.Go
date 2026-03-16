// Problem 22: Names Scores
// Total of all name scores in the file.
// Answer: 871198282

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func nameValue(name string) int {
	sum := 0
	for _, ch := range name {
		sum += int(ch-'A') + 1
	}
	return sum
}

func loadNames() []string {
	data, err := os.ReadFile("names.txt")
	if err != nil {
		fmt.Println("Error reading names.txt:", err)
		os.Exit(1)
	}

	raw := strings.ReplaceAll(string(data), "\"", "")
	return strings.Split(raw, ",")
}

var names []string

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

func benchmark(iterations int) time.Duration {
	// Warmup
	for i := 0; i < 10; i++ {
		solve()
	}

	start := time.Now()
	var result int64
	for i := 0; i < iterations; i++ {
		result = solve()
	}
	elapsed := time.Since(start)
	fmt.Printf("Result: %d (%.2f ns/op)\n", result, float64(elapsed.Nanoseconds())/float64(iterations))
	return elapsed
}

func main() {
	const iterations = 1000

	fmt.Println("Problem 22: Names Scores")
	fmt.Println("========================")
	fmt.Printf("Computing total of all name scores, Iterations: %d\n\n", iterations)

	names = loadNames()
	benchmark(iterations)
}
