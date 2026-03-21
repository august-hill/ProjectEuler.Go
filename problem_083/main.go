// Problem 83: Path Sum: Four Ways
// Find the minimal path sum from top-left to bottom-right, moving in all 4 directions.
// Answer: 425185

package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

var (
	matrix083  [80][80]int
	rows083    int
	cols083    int
	loadOnce083 sync.Once
)

func loadData083() {
	loadOnce083.Do(func() {
		data, err := os.ReadFile("p083_matrix.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not open p083_matrix.txt\n")
			os.Exit(1)
		}

		lines := strings.Split(strings.TrimSpace(string(data)), "\n")
		rows083 = 0
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if len(line) < 2 {
				continue
			}
			cols083 = 0
			for _, tok := range strings.Split(line, ",") {
				v, _ := strconv.Atoi(strings.TrimSpace(tok))
				matrix083[rows083][cols083] = v
				cols083++
			}
			rows083++
		}
	})
}

// Min-heap for Dijkstra
type node083 struct {
	cost int
	r, c int
}

type minHeap083 []node083

func (h minHeap083) Len() int            { return len(h) }
func (h minHeap083) Less(i, j int) bool  { return h[i].cost < h[j].cost }
func (h minHeap083) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap083) Push(x interface{}) { *h = append(*h, x.(node083)) }
func (h *minHeap083) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func solve() int64 {
	loadData083()

	dist := [80][80]int{}
	for i := 0; i < rows083; i++ {
		for j := 0; j < cols083; j++ {
			dist[i][j] = math.MaxInt64
		}
	}

	h := &minHeap083{}
	heap.Init(h)

	dist[0][0] = matrix083[0][0]
	heap.Push(h, node083{matrix083[0][0], 0, 0})

	dr := [4]int{-1, 1, 0, 0}
	dc := [4]int{0, 0, -1, 1}

	for h.Len() > 0 {
		cur := heap.Pop(h).(node083)
		if cur.cost > dist[cur.r][cur.c] {
			continue
		}
		if cur.r == rows083-1 && cur.c == cols083-1 {
			return int64(cur.cost)
		}

		for d := 0; d < 4; d++ {
			nr := cur.r + dr[d]
			nc := cur.c + dc[d]
			if nr >= 0 && nr < rows083 && nc >= 0 && nc < cols083 {
				newCost := cur.cost + matrix083[nr][nc]
				if newCost < dist[nr][nc] {
					dist[nr][nc] = newCost
					heap.Push(h, node083{newCost, nr, nc})
				}
			}
		}
	}

	return int64(dist[rows083-1][cols083-1])
}

func main() { bench.Run(83, solve) }
