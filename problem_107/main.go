// Problem 107: Minimal Network
// Find the maximum saving by removing redundant edges using Kruskal's MST.
// Answer: 259679

package main

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

//go:embed p107_network.txt
var networkData string

type edge107 struct{ u, v, w int }

var parent107 [40]int

func find107(x int) int {
	for parent107[x] != x {
		parent107[x] = parent107[parent107[x]]
		x = parent107[x]
	}
	return x
}

func unite107(a, b int) bool {
	a = find107(a)
	b = find107(b)
	if a == b {
		return false
	}
	parent107[a] = b
	return true
}

func solve() int64 {
	lines := strings.Split(strings.TrimSpace(networkData), "\n")
	n := len(lines)

	adj := make([][]int, n)
	for i := range adj {
		adj[i] = make([]int, n)
	}

	for i, line := range lines {
		line = strings.TrimSpace(line)
		parts := strings.Split(line, ",")
		for j, p := range parts {
			p = strings.TrimSpace(p)
			if p != "-" {
				v, _ := strconv.Atoi(p)
				adj[i][j] = v
			}
		}
	}

	totalWeight := 0
	var edges []edge107
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if adj[i][j] > 0 {
				totalWeight += adj[i][j]
				edges = append(edges, edge107{i, j, adj[i][j]})
			}
		}
	}

	sort.Slice(edges, func(a, b int) bool {
		return edges[a].w < edges[b].w
	})

	for i := 0; i < n; i++ {
		parent107[i] = i
	}

	mstWeight := 0
	for _, e := range edges {
		if unite107(e.u, e.v) {
			mstWeight += e.w
		}
	}

	return int64(totalWeight - mstWeight)
}

func main() { bench.Run(107, solve) }
