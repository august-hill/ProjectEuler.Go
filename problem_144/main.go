// Problem 144: Investigating Multiple Reflections of a Laser Beam
// How many reflections occur before the beam exits?
// Answer: 354

package main

import (
	"math"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

func solve() int64 {
	x0, y0 := 0.0, 10.1
	x1, y1 := 1.4, -9.6
	count := 0

	for {
		// Normal at (x1, y1): proportional to (4*x1, y1)
		nx, ny := 4.0*x1, y1

		// Incoming direction
		dx, dy := x1-x0, y1-y0

		// Reflect: d' = d - 2*(d.n / n.n)*n
		dot := dx*nx + dy*ny
		nn := nx*nx + ny*ny
		rx := dx - 2.0*dot/nn*nx
		ry := dy - 2.0*dot/nn*ny

		// Next intersection: 4(x1+t*rx)^2 + (y1+t*ry)^2 = 100
		a := 4.0*rx*rx + ry*ry
		b := 8.0*x1*rx + 2.0*y1*ry
		c := 4.0*x1*x1 + y1*y1 - 100.0

		disc := b*b - 4.0*a*c
		t := (-b + math.Sqrt(disc)) / (2.0 * a)
		if math.Abs(t) < 1e-9 {
			t = (-b - math.Sqrt(disc)) / (2.0 * a)
		}

		x0, y0 = x1, y1
		x1 = x0 + t*rx
		y1 = y0 + t*ry
		count++

		if y1 > 0 && math.Abs(x1) <= 0.01 {
			break
		}
	}
	return int64(count)
}

func main() { bench.Run(144, solve) }
