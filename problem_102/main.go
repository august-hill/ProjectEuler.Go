// Problem 102: Triangle Containment
// How many triangles contain the origin?
// Answer: 228

package main

import (
	_ "embed"
	"strings"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

//go:embed p102_triangles.txt
var triangleData string

func containsOrigin102(x1, y1, x2, y2, x3, y3 int) bool {
	d1 := int64(-x1)*int64(y2-y1) - int64(-y1)*int64(x2-x1)
	d2 := int64(-x2)*int64(y3-y2) - int64(-y2)*int64(x3-x2)
	d3 := int64(-x3)*int64(y1-y3) - int64(-y3)*int64(x1-x3)

	hasNeg := d1 < 0 || d2 < 0 || d3 < 0
	hasPos := d1 > 0 || d2 > 0 || d3 > 0
	return !(hasNeg && hasPos)
}

func solve() int64 {
	count := 0
	for _, line := range strings.Split(strings.TrimSpace(triangleData), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		var x1, y1, x2, y2, x3, y3 int
		n, _ := parseInts102(line, &x1, &y1, &x2, &y2, &x3, &y3)
		if n == 6 && containsOrigin102(x1, y1, x2, y2, x3, y3) {
			count++
		}
	}
	return int64(count)
}

func parseInts102(s string, vals ...*int) (int, error) {
	parts := strings.Split(s, ",")
	count := 0
	for i, p := range parts {
		if i >= len(vals) {
			break
		}
		p = strings.TrimSpace(p)
		var v int
		neg := false
		if len(p) > 0 && p[0] == '-' {
			neg = true
			p = p[1:]
		}
		for _, c := range p {
			v = v*10 + int(c-'0')
		}
		if neg {
			v = -v
		}
		*vals[i] = v
		count++
	}
	return count, nil
}

func main() { bench.Run(102, solve) }
