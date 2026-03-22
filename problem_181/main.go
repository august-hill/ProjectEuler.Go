// Problem 181: Investigating in How Many Ways Objects of Two Different Colours Can Be Grouped
// Answer: 83735848679360680

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

const B181 = 60
const W181 = 40

func solve() int64 {
	var dp [B181 + 1][W181 + 1]int64
	dp[B181][W181] = 1

	for b := 0; b <= B181; b++ {
		for w := 0; w <= W181; w++ {
			if b == 0 && w == 0 {
				continue
			}
			var ndp [B181 + 1][W181 + 1]int64
			for rb := 0; rb <= B181; rb++ {
				for rw := 0; rw <= W181; rw++ {
					if dp[rb][rw] == 0 {
						continue
					}
					for k := 0; ; k++ {
						nb := rb - k*b
						nw := rw - k*w
						if nb < 0 || nw < 0 {
							break
						}
						ndp[nb][nw] += dp[rb][rw]
					}
				}
			}
			dp = ndp
		}
	}
	return dp[0][0]
}

func main() { bench.Run(181, solve) }
