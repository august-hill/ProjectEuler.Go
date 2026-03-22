// Problem 184: Triangles Containing the Origin
// Answer: 1725323624056

package main

import (
	"sort"
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const R184 = 105

type pt184 struct{ x, y int }

func ptQuadrant184(p pt184) int {
	if p.y > 0 {
		if p.x >= 0 {
			return 0
		}
		return 1
	}
	if p.y < 0 {
		if p.x <= 0 {
			return 2
		}
		return 3
	}
	if p.x > 0 {
		return 0
	}
	return 2
}

var (
	once184        sync.Once
	answerCache184 int64
)

func compute184() {
	r2 := R184 * R184

	var pts []pt184
	for x := -(R184 - 1); x <= R184-1; x++ {
		for y := -(R184 - 1); y <= R184-1; y++ {
			if x*x+y*y < r2 && (x != 0 || y != 0) {
				pts = append(pts, pt184{x, y})
			}
		}
	}

	sort.Slice(pts, func(i, j int) bool {
		p, q := pts[i], pts[j]
		qp, qq := ptQuadrant184(p), ptQuadrant184(q)
		if qp != qq {
			return qp < qq
		}
		cross := int64(p.x)*int64(q.y) - int64(p.y)*int64(q.x)
		return cross > 0
	})

	n := len(pts)
	bad := int64(0)
	j := 1
	for i := 0; i < n; i++ {
		if j <= i {
			j = i + 1
		}
		for j < i+n {
			jj := j % n
			cross := int64(pts[i].x)*int64(pts[jj].y) - int64(pts[i].y)*int64(pts[jj].x)
			if cross < 0 {
				break
			}
			if cross == 0 {
				dot := int64(pts[i].x)*int64(pts[jj].x) + int64(pts[i].y)*int64(pts[jj].y)
				if dot <= 0 {
					break
				}
			}
			j++
		}
		fi := int64(j - i - 1)
		bad += fi * (fi - 1) / 2
	}

	totalTri := int64(n) * int64(n-1) * int64(n-2) / 6
	answerCache184 = totalTri - bad
}

func solve() int64 {
	once184.Do(compute184)
	return answerCache184
}

func main() { bench.Run(184, solve) }
