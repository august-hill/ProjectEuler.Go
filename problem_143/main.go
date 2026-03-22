// Problem 143: Torricelli Triangles
// Find the sum of all Torricelli triangle perimeters below 120000.
// Answer: 30758397

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

const limit143 = 120000

func gcd143(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

const hashSize143 = 1 << 22
const hashMask143 = hashSize143 - 1

var pairKeys143 [hashSize143][2]int
var pairUsed143 [hashSize143]bool

func pairHash143(x, y int) int {
	h := uint(x)*1000003 + uint(y)
	return int(h & hashMask143)
}

func setPair143(x, y int) {
	h := pairHash143(x, y)
	for pairUsed143[h] {
		if pairKeys143[h][0] == x && pairKeys143[h][1] == y {
			return
		}
		h = (h + 1) & hashMask143
	}
	pairKeys143[h][0] = x
	pairKeys143[h][1] = y
	pairUsed143[h] = true
}

func hasPair143(x, y int) bool {
	h := pairHash143(x, y)
	for pairUsed143[h] {
		if pairKeys143[h][0] == x && pairKeys143[h][1] == y {
			return true
		}
		h = (h + 1) & hashMask143
	}
	return false
}

func solve() int64 {
	// Build adjacency for pairs (x, y) satisfying x^2 + xy + y^2 = z^2
	adjSize := make([]int, limit143+1)
	type pair struct{ a, b int }
	var allPairs []pair

	for m := 2; m <= 500; m++ {
		for n := 1; n < m; n++ {
			if gcd143(m, n) != 1 {
				continue
			}
			if (m-n)%3 == 0 {
				continue
			}
			x := m*m - n*n
			y := 2*m*n + n*n

			for k := 1; ; k++ {
				kx, ky := k*x, k*y
				if kx+ky > limit143 {
					break
				}
				setPair143(kx, ky)
				setPair143(ky, kx)
				allPairs = append(allPairs, pair{kx, ky})
				if kx != ky {
					allPairs = append(allPairs, pair{ky, kx})
				}
			}
		}
	}

	for _, p := range allPairs {
		if p.a < limit143+1 {
			adjSize[p.a]++
		}
	}

	adjStart := make([]int, limit143+2)
	adjStart[0] = 0
	for i := 1; i <= limit143+1; i++ {
		if i <= limit143 {
			adjStart[i] = adjStart[i-1] + adjSize[i-1]
		} else {
			adjStart[i] = adjStart[i-1]
		}
	}

	totalPairs := adjStart[limit143+1]
	allAdj := make([]int, totalPairs)
	fill := make([]int, limit143+1)
	for _, p := range allPairs {
		v := p.a
		if v <= limit143 {
			allAdj[adjStart[v]+fill[v]] = p.b
			fill[v]++
		}
	}

	seen := make([]bool, limit143+1)
	var total int64

	for p := 1; p <= limit143; p++ {
		if adjSize[p] == 0 {
			continue
		}
		pAdj := allAdj[adjStart[p] : adjStart[p]+adjSize[p]]

		for _, q := range pAdj {
			if q <= p {
				continue
			}
			maxR := limit143 - p - q
			if maxR < 1 {
				continue
			}

			qAdj := allAdj[adjStart[q] : adjStart[q]+adjSize[q]]
			for _, r := range qAdj {
				if r > maxR {
					continue
				}
				if hasPair143(p, r) {
					s := p + q + r
					if !seen[s] {
						seen[s] = true
						total += int64(s)
					}
				}
			}
		}
	}

	return total
}

func main() { bench.Run(143, solve) }
