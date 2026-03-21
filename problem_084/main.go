// Problem 84: Monopoly Odds
// Using two 4-sided dice, find the three most popular squares on a Monopoly board.
// Answer: 101524

package main

import (
	"github.com/august-hill/ProjectEuler.Go/bench"
)

const (
	goSquare084 = 0
	jail084     = 10
	g2j084      = 30
	cc1_084     = 2
	cc2_084     = 17
	cc3_084     = 33
	ch1_084     = 7
	ch2_084     = 22
	ch3_084     = 36
	c1_084      = 11
	e3_084      = 24
	h2_084      = 39
	r1_084      = 5
)

func applyLanding084(freq []float64, square int, prob float64) {
	switch {
	case square == g2j084:
		freq[jail084] += prob
	case square == cc1_084 || square == cc2_084 || square == cc3_084:
		freq[goSquare084] += prob / 16.0
		freq[jail084] += prob / 16.0
		freq[square] += prob * 14.0 / 16.0
	case square == ch1_084 || square == ch2_084 || square == ch3_084:
		var nextR, nextU int
		switch square {
		case ch1_084:
			nextR, nextU = 15, 12
		case ch2_084:
			nextR, nextU = 25, 28
		default:
			nextR, nextU = 5, 12
		}
		back3 := (square + 40 - 3) % 40

		freq[goSquare084] += prob / 16.0
		freq[jail084] += prob / 16.0
		freq[c1_084] += prob / 16.0
		freq[e3_084] += prob / 16.0
		freq[h2_084] += prob / 16.0
		freq[r1_084] += prob / 16.0
		freq[nextR] += prob * 2.0 / 16.0
		freq[nextU] += prob / 16.0
		applyLanding084(freq, back3, prob/16.0)
		freq[square] += prob * 6.0 / 16.0
	default:
		freq[square] += prob
	}
}

func solve() int64 {
	diceProb := make([]float64, 9)
	for d1 := 1; d1 <= 4; d1++ {
		for d2 := 1; d2 <= 4; d2++ {
			diceProb[d1+d2] += 1.0 / 16.0
		}
	}

	freq := make([]float64, 40)
	freq[0] = 1.0

	for iter := 0; iter < 200; iter++ {
		newFreq := make([]float64, 40)
		for pos := 0; pos < 40; pos++ {
			if freq[pos] == 0.0 {
				continue
			}
			// 3 consecutive doubles -> jail
			newFreq[jail084] += freq[pos] * (1.0 / 64.0)
			remaining := freq[pos] * (63.0 / 64.0)
			for sum := 2; sum <= 8; sum++ {
				next := (pos + sum) % 40
				applyLanding084(newFreq, next, remaining*diceProb[sum])
			}
		}
		copy(freq, newFreq)
	}

	// Find top 3
	top := [3]int{-1, -1, -1}
	for t := 0; t < 3; t++ {
		best := -1
		for i := 0; i < 40; i++ {
			if i == top[0] || i == top[1] {
				continue
			}
			if best == -1 || freq[i] > freq[best] {
				best = i
			}
		}
		top[t] = best
	}

	return int64(top[0]*10000 + top[1]*100 + top[2])
}

func main() { bench.Run(84, solve) }
