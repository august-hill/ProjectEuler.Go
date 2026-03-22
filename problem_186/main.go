// Problem 186: Connectedness of a Network
// Answer: 2325629

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

const users186 = 1000000

var (
	parent186 [users186]int
	sz186     [users186]int
)

func ufInit186() {
	for i := 0; i < users186; i++ {
		parent186[i] = i
		sz186[i] = 1
	}
}

func ufFind186(x int) int {
	for parent186[x] != x {
		parent186[x] = parent186[parent186[x]]
		x = parent186[x]
	}
	return x
}

func ufUnion186(a, b int) {
	a, b = ufFind186(a), ufFind186(b)
	if a == b {
		return
	}
	if sz186[a] < sz186[b] {
		a, b = b, a
	}
	parent186[b] = a
	sz186[a] += sz186[b]
}

func solve() int64 {
	var rb [55]int64
	for k := 1; k <= 55; k++ {
		val := (100003 - 200003*int64(k) + 300007*int64(k)*int64(k)*int64(k)) % 1000000
		if val < 0 {
			val += 1000000
		}
		rb[k-1] = val
	}

	ufInit186()
	PM := 524287
	target := 990000

	consumed := 0
	rp := 0
	calls := 0

	for {
		genNext := func() int64 {
			var val int64
			if consumed < 55 {
				val = rb[consumed]
			} else {
				i24 := (rp + 55 - 24) % 55
				i55 := rp
				val = (rb[i24] + rb[i55]) % 1000000
				rb[rp] = val
				rp = (rp + 1) % 55
			}
			consumed++
			return val
		}

		callerVal := genNext()
		calledVal := genNext()

		caller := int(callerVal)
		called := int(calledVal)

		if caller == called {
			continue
		}

		calls++
		ufUnion186(caller, called)

		if sz186[ufFind186(PM)] >= target {
			return int64(calls)
		}
	}
}

func main() { bench.Run(186, solve) }
