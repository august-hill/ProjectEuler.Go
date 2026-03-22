// Problem 126: Cuboid Layers
// Find the smallest number of cubes that cannot be achieved by exactly 1000 cuboids.
// Answer: 18522

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const limit126 = 20000

func layerCubes126(a, b, c, k int) int64 {
	return 2*int64(a*b+b*c+a*c) +
		4*int64(k-1)*int64(a+b+c) +
		4*int64(k-1)*int64(k-2)
}

func solve() int64 {
	count := make([]int, limit126+1)

	for a := 1; a <= limit126; a++ {
		for b := a; ; b++ {
			if layerCubes126(a, b, b, 1) > limit126 {
				break
			}
			for c := b; ; c++ {
				if layerCubes126(a, b, c, 1) > limit126 {
					break
				}
				for k := 1; ; k++ {
					cubes := layerCubes126(a, b, c, k)
					if cubes > limit126 {
						break
					}
					count[cubes]++
				}
			}
		}
	}

	for n := 1; n <= limit126; n++ {
		if count[n] == 1000 {
			return int64(n)
		}
	}
	return -1
}

func main() { bench.Run(126, solve) }
