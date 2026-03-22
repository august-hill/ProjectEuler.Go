// Problem 161: Triominoes
// Answer: 20574308184277971

package main

import (
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const r161 = 12
const c161 = 9
const ncells161 = r161 * c161
const window161 = 2*c161 + 1

var shapes161 = [6][3][2]int{
	{{0, 0}, {0, 1}, {0, 2}},
	{{0, 0}, {1, 0}, {2, 0}},
	{{0, 0}, {0, 1}, {1, 0}},
	{{0, 0}, {0, 1}, {1, 1}},
	{{0, 0}, {1, 0}, {1, 1}},
	{{0, 0}, {1, -1}, {1, 0}},
}

var (
	once161        sync.Once
	answerCache161 int64
)

func compute161() {
	curDp := make([]int64, 1<<window161)
	nxtDp := make([]int64, 1<<window161)

	curDp[0] = 1

	for pos := 0; pos < ncells161; pos++ {
		row := pos / c161
		col := pos % c161

		for i := range nxtDp {
			nxtDp[i] = 0
		}

		for mask := 0; mask < (1 << window161); mask++ {
			if curDp[mask] == 0 {
				continue
			}
			ways := curDp[mask]

			if mask&1 != 0 {
				nxtDp[mask>>1] += ways
				continue
			}

			for s := 0; s < 6; s++ {
				ok := true
				var offsets [3]int
				for k := 0; k < 3; k++ {
					r2 := row + shapes161[s][k][0]
					c2 := col + shapes161[s][k][1]
					if r2 < 0 || r2 >= r161 || c2 < 0 || c2 >= c161 {
						ok = false
						break
					}
					pos2 := r2*c161 + c2
					off := pos2 - pos
					if off < 0 || off >= window161 {
						ok = false
						break
					}
					offsets[k] = off
				}
				if !ok {
					continue
				}

				newMask := mask
				conflict := false
				for k := 0; k < 3; k++ {
					if newMask&(1<<offsets[k]) != 0 {
						conflict = true
						break
					}
					newMask |= (1 << offsets[k])
				}
				if conflict {
					continue
				}
				nxtDp[newMask>>1] += ways
			}
		}
		curDp, nxtDp = nxtDp, curDp
	}

	answerCache161 = curDp[0]
}

func solve() int64 {
	once161.Do(compute161)
	return answerCache161
}

func main() { bench.Run(161, solve) }
