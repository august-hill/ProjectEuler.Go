// Problem 165: Intersections
// Answer: 2868868

package main

import (
	"sort"
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const nSeg165 = 5000

type point165 struct{ x, y int64 }
type ipoint165 struct{ px, py, pq int64 }

var (
	once165        sync.Once
	answerCache165 int64
)

func gcd165(a, b int64) int64 {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs165(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func compute165() {
	s := int64(290797)
	t := make([]int64, 4*nSeg165)
	for i := range t {
		s = (s * s) % 50515093
		t[i] = s % 500
	}

	seg := [nSeg165][2]point165{}
	for i := 0; i < nSeg165; i++ {
		seg[i][0] = point165{t[4*i], t[4*i+1]}
		seg[i][1] = point165{t[4*i+2], t[4*i+3]}
	}

	ipoints := make([]ipoint165, 0, 1<<20)

	for i := 0; i < nSeg165; i++ {
		ax, ay := seg[i][0].x, seg[i][0].y
		bx, by := seg[i][1].x, seg[i][1].y
		dx, dy := bx-ax, by-ay

		for j := i + 1; j < nSeg165; j++ {
			cx, cy := seg[j][0].x, seg[j][0].y
			ex, ey := seg[j][1].x, seg[j][1].y
			fx, fy := ex-cx, ey-cy

			denom := dx*fy - dy*fx
			if denom == 0 {
				continue
			}

			tNum := (cx-ax)*fy - (cy-ay)*fx
			uNum := (cx-ax)*dy - (cy-ay)*dx

			if denom > 0 {
				if tNum <= 0 || tNum >= denom {
					continue
				}
				if uNum <= 0 || uNum >= denom {
					continue
				}
			} else {
				if tNum >= 0 || tNum <= denom {
					continue
				}
				if uNum >= 0 || uNum <= denom {
					continue
				}
			}

			px := ax*denom + tNum*dx
			py := ay*denom + tNum*dy
			pq := denom

			if pq < 0 {
				px, py, pq = -px, -py, -pq
			}
			g := gcd165(gcd165(abs165(px), abs165(py)), pq)
			if g > 0 {
				px /= g
				py /= g
				pq /= g
			}
			ipoints = append(ipoints, ipoint165{px, py, pq})
		}
	}

	sort.Slice(ipoints, func(i, j int) bool {
		a, b := ipoints[i], ipoints[j]
		if a.pq != b.pq {
			return a.pq < b.pq
		}
		if a.px != b.px {
			return a.px < b.px
		}
		return a.py < b.py
	})

	count := int64(0)
	for i := 0; i < len(ipoints); {
		j := i + 1
		for j < len(ipoints) && ipoints[j] == ipoints[i] {
			j++
		}
		count++
		i = j
	}

	answerCache165 = count
}

func solve() int64 {
	once165.Do(compute165)
	return answerCache165
}

func main() { bench.Run(165, solve) }
