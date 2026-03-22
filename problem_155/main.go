// Problem 155: Counting Capacitor Circuits
// Answer: 3857447

package main

import (
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const maxN155 = 19
const maxFracs155 = 2000000
const hashSize155 = 1 << 24
const hashMask155 = hashSize155 - 1

type frac155 struct{ p, q int64 }

type entry155 struct {
	p, q int64
	next int
}

var (
	once155       sync.Once
	result155     int64
	hashHeads155  [hashSize155]int
	entryPool155  []entry155
	poolIdx155    int
	totalDistinct155 int
	levelFracs155 [maxN155][]frac155
	levelCount155 [maxN155]int
)

func gcd155(a, b int64) int64 {
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

func hashFrac155(p, q int64) int {
	h := uint32(p*1000003 + q*999983)
	h ^= uint32(p >> 16)
	h ^= uint32(q >> 16)
	return int(h) & hashMask155
}

func hashInsert155(p, q int64) bool {
	h := hashFrac155(p, q)
	for idx := hashHeads155[h]; idx >= 0; idx = entryPool155[idx].next {
		if entryPool155[idx].p == p && entryPool155[idx].q == q {
			return false
		}
	}
	entryPool155[poolIdx155] = entry155{p, q, hashHeads155[h]}
	hashHeads155[h] = poolIdx155
	poolIdx155++
	totalDistinct155++
	return true
}

func compute155() {
	entryPool155 = make([]entry155, 10000000)
	for i := range hashHeads155 {
		hashHeads155[i] = -1
	}
	poolIdx155 = 0
	totalDistinct155 = 0

	for i := 1; i < maxN155; i++ {
		levelFracs155[i] = make([]frac155, maxFracs155)
		levelCount155[i] = 0
	}

	levelFracs155[1][0] = frac155{1, 1}
	levelCount155[1] = 1
	hashInsert155(1, 1)

	for n := 2; n <= 18; n++ {
		newCount := 0
		for k := 1; k <= n/2; k++ {
			j := n - k
			for a := 0; a < levelCount155[k]; a++ {
				ap := levelFracs155[k][a].p
				aq := levelFracs155[k][a].q
				bStart := 0
				if k == j {
					bStart = a
				}
				for b := bStart; b < levelCount155[j]; b++ {
					bp := levelFracs155[j][b].p
					bq := levelFracs155[j][b].q

					// Parallel: a + b
					pp := ap*bq + bp*aq
					pq := aq * bq
					g := gcd155(pp, pq)
					pp /= g
					pq /= g
					if hashInsert155(pp, pq) && newCount < maxFracs155 {
						levelFracs155[n][newCount] = frac155{pp, pq}
						newCount++
					}

					// Series: ab/(a+b)
					sp := ap * bp
					sq := ap*bq + bp*aq
					g = gcd155(sp, sq)
					sp /= g
					sq /= g
					if hashInsert155(sp, sq) && newCount < maxFracs155 {
						levelFracs155[n][newCount] = frac155{sp, sq}
						newCount++
					}
				}
			}
		}
		levelCount155[n] = newCount
	}

	result155 = int64(totalDistinct155)
}

func solve() int64 {
	once155.Do(compute155)
	return result155
}

func main() { bench.Run(155, solve) }
