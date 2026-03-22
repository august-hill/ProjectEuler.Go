// Problem 192: Best Approximations
// Answer: 57060635927998347

package main

import (
	"math"
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const bound192 = 1000000000000

var (
	once192        sync.Once
	answerCache192 int64
)

func isPerfectSquare192(n int) bool {
	s := int(math.Sqrt(float64(n)))
	for s*s > n {
		s--
	}
	for (s+1)*(s+1) <= n {
		s++
	}
	return s*s == n
}

// Compare |p1/q1 - sqrt(n)| vs |p2/q2 - sqrt(n)|. Returns true if p1/q1 is closer.
func closerToSqrt192(p1, q1, p2, q2 int64, n int) bool {
	sq := math.Sqrt(float64(n))
	e1 := math.Abs(float64(p1)/float64(q1) - sq)
	e2 := math.Abs(float64(p2)/float64(q2) - sq)
	if e1 < e2*(1.0-1e-15) {
		return true
	}
	if e2 < e1*(1.0-1e-15) {
		return false
	}
	// Approximate comparison for very close values
	lhs := float64(p1*p1-int64(n)*q1*q1) * float64(q2)
	if lhs < 0 {
		lhs = -lhs
	}
	rhs := float64(p2*p2-int64(n)*q2*q2) * float64(q1)
	if rhs < 0 {
		rhs = -rhs
	}
	// p1/q1 closer means smaller error: e1i/q1*(p1+q1*sq) < e2i/q2*(p2+q2*sq)
	lhsFull := lhs * (float64(p2) + float64(q2)*sq)
	rhsFull := rhs * (float64(p1) + float64(q1)*sq)
	return lhsFull < rhsFull
}

func bestDenom192(n int) int64 {
	a0 := int(math.Sqrt(float64(n)))
	for (a0+1)*(a0+1) <= n {
		a0++
	}
	for a0*a0 > n {
		a0--
	}

	pp, pq := int64(1), int64(0)
	cp, cq := int64(a0), int64(1)

	m, d, a := int64(0), int64(1), int64(a0)

	for {
		m = d*a - m
		d = (int64(n) - m*m) / d
		if d == 0 {
			break
		}
		a = (int64(a0) + m) / d

		np := a*cp + pp
		nq := a*cq + pq

		if nq > bound192 {
			jMax := int64(0)
			if pq >= 0 && cq > 0 {
				jMax = (bound192 - pq) / cq
			}
			if jMax >= 1 {
				sq := jMax*cq + pq
				sp := jMax*cp + pp
				if closerToSqrt192(sp, sq, cp, cq, n) {
					return sq
				}
			}
			return cq
		}

		pp, pq = cp, cq
		cp, cq = np, nq
	}
	return cq
}

func compute192() {
	total := int64(0)
	for n := 2; n <= 100000; n++ {
		if isPerfectSquare192(n) {
			continue
		}
		total += bestDenom192(n)
	}
	answerCache192 = total
}

func solve() int64 {
	once192.Do(compute192)
	return answerCache192
}

func main() { bench.Run(192, solve) }
