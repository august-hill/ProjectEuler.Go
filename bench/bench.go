// Package bench provides a shared benchmark utility for Project Euler Go solutions.
//
// Usage:
//
//	import "github.com/august-hill/ProjectEuler.Go/bench"
//	func solve() int64 { /* return answer */ }
//	func main() { bench.Run(60, solve) }
package bench

import (
	"fmt"
	"sort"
	"time"
)

// sink prevents the compiler from optimizing away solve calls
var sink int64

// Run benchmarks the given solve function and prints the standard BENCHMARK line.
func Run(problem int, solve func() int64) {
	// Warmup: 3 runs
	for i := 0; i < 3; i++ {
		sink = solve()
	}

	// Calibrate: time one run
	t0 := time.Now()
	sink = solve()
	calNs := time.Since(t0).Nanoseconds()

	var iters int
	switch {
	case calNs < 1_000_000:
		iters = 1000
	case calNs < 100_000_000:
		iters = 100
	case calNs < 1_000_000_000:
		iters = 10
	default:
		iters = 3
	}

	// Timed runs
	times := make([]int64, iters)
	var answer int64
	for i := 0; i < iters; i++ {
		t0 = time.Now()
		answer = solve()
		times[i] = time.Since(t0).Nanoseconds()
	}

	// Median
	sort.Slice(times, func(i, j int) bool { return times[i] < times[j] })
	medianNs := times[iters/2]

	fmt.Printf("BENCHMARK|problem=%03d|answer=%d|time_ns=%d|iterations=%d\n",
		problem, answer, medianNs, iters)
}
