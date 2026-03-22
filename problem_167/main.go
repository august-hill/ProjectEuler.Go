// Problem 167: Investigating Ulam sequences
// Answer: 3916160068885

package main

import (
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const target167 = 100000000000

var (
	once167        sync.Once
	answerCache167 int64
)

func cntGet167(a []byte, i int64) int {
	return int((a[i>>2] >> ((i & 3) << 1)) & 3)
}

func cntInc167(a []byte, i int64) {
	shift := (i & 3) << 1
	v := int((a[i>>2] >> shift) & 3)
	if v < 2 {
		a[i>>2] = byte((int(a[i>>2]) & ^(3 << shift)) | ((v + 1) << shift))
	}
}

func ulamNth167(b int, n int64) int64 {
	maxTerms := 500000
	if b >= 15 {
		maxTerms = 1000000
	}

	seq := make([]int64, maxTerms)
	diffs := make([]int64, maxTerms)

	seq[0] = 2
	seq[1] = int64(b)
	count := 2

	maxVal := int64(maxTerms) * int64(b+2)
	if maxVal < 500000 {
		maxVal = 500000
	}
	arrBytes := maxVal/4 + 2
	cnt := make([]byte, arrBytes)
	cntInc167(cnt, 2+int64(b))

	ndiffs := 0
	period := -1
	periodDstart := -1

outer:
	for count < maxTerms {
		prev := seq[count-1]

		if prev*2+int64(b)+1000 >= maxVal {
			nmax := maxVal + int64(maxTerms)*int64(b) + 1000000
			narr := nmax/4 + 2
			newCnt := make([]byte, narr)
			copy(newCnt, cnt)
			cnt = newCnt
			arrBytes = narr
			maxVal = nmax
		}

		nextVal := int64(-1)
		for v := prev + 1; v <= maxVal; v++ {
			if cntGet167(cnt, v) == 1 {
				nextVal = v
				break
			}
		}
		if nextVal == -1 {
			break
		}

		seq[count] = nextVal
		diffs[ndiffs] = nextVal - prev
		ndiffs++

		for i := 0; i < count; i++ {
			sum := seq[i] + nextVal
			if sum < maxVal {
				cntInc167(cnt, sum)
			}
		}
		count++

		if period == -1 && ndiffs >= 2000 && ndiffs%1000 == 0 {
			for P := 1; P <= ndiffs/25; P++ {
				if 25*P > ndiffs {
					break
				}
				match := true
				for k := 0; k < 20*P; k++ {
					if diffs[ndiffs-1-k] != diffs[ndiffs-1-k-P] {
						match = false
						break
					}
				}
				if match {
					period = P
					periodDstart = 0
					for i := ndiffs - 2*P - 1; i >= 1; i-- {
						if diffs[i] != diffs[i+P] {
							periodDstart = i + 1
							break
						}
					}
					break outer
				}
			}
		}
	}

	if period == -1 {
		if n <= int64(count) {
			return seq[n-1]
		}
		return -1
	}

	ps := int64(0)
	for i := 0; i < period; i++ {
		ps += diffs[periodDstart+i]
	}
	s01 := int64(periodDstart) + 2
	s0v := seq[periodDstart+1]

	if n < s01 {
		return seq[n-1]
	}

	off := n - s01
	fp := off / int64(period)
	rem := off % int64(period)
	bv := s0v
	for k := int64(0); k < rem; k++ {
		bv += diffs[periodDstart+int(k)]
	}
	return bv + fp*ps
}

func compute167() {
	total := int64(0)
	for k := 5; k <= 21; k += 2 {
		total += ulamNth167(k, target167)
	}
	answerCache167 = total
}

func solve() int64 {
	once167.Do(compute167)
	return answerCache167
}

func main() { bench.Run(167, solve) }
