// Problem 185: Number Mind
// Answer: 4640261571849533

package main

import (
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const N185 = 16
const NClues185 = 22

var (
	once185      sync.Once
	clueDigits185 [NClues185][N185]int
	clueCorrect185 [NClues185]int
	secret185    [N185]int
	found185     bool
	answer185    int64
)

func initClues185() {
	gs := []string{
		"5616185650518293", "3847439647293047", "5855462940810587",
		"9742855507068353", "4296849643607543", "3174248439465858",
		"4513559094146117", "7890971548908067", "8157356344118483",
		"2615250744386899", "8690095851526254", "6375711915077050",
		"6913859173121360", "6442889055042768", "2321386104303845",
		"2326509471271448", "5251583379644322", "1748270476758276",
		"4895722652190306", "3041631117224635", "1841236454324589",
		"2659862637316867",
	}
	cc := []int{2, 1, 3, 3, 3, 1, 2, 3, 1, 2, 3, 1, 1, 2, 0, 2, 2, 3, 1, 3, 3, 2}
	for i := 0; i < NClues185; i++ {
		for j := 0; j < N185; j++ {
			clueDigits185[i][j] = int(gs[i][j] - '0')
		}
		clueCorrect185[i] = cc[i]
	}
}

func checkPartial185(pos int) bool {
	for c := 0; c < NClues185; c++ {
		matches := 0
		for i := 0; i < pos; i++ {
			if secret185[i] == clueDigits185[c][i] {
				matches++
			}
		}
		remaining := N185 - pos
		if matches > clueCorrect185[c] {
			return false
		}
		if matches+remaining < clueCorrect185[c] {
			return false
		}
	}
	return true
}

func backtrack185(pos int) {
	if found185 {
		return
	}
	if pos == N185 {
		for c := 0; c < NClues185; c++ {
			matches := 0
			for i := 0; i < N185; i++ {
				if secret185[i] == clueDigits185[c][i] {
					matches++
				}
			}
			if matches != clueCorrect185[c] {
				return
			}
		}
		found185 = true
		answer185 = 0
		for i := 0; i < N185; i++ {
			answer185 = answer185*10 + int64(secret185[i])
		}
		return
	}

	for d := 0; d <= 9; d++ {
		valid := true
		for c := 0; c < NClues185; c++ {
			if clueCorrect185[c] == 0 && clueDigits185[c][pos] == d {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}
		secret185[pos] = d
		if checkPartial185(pos + 1) {
			backtrack185(pos + 1)
		}
		if found185 {
			return
		}
	}
}

func compute185() {
	initClues185()
	found185 = false
	answer185 = 0
	for i := range secret185 {
		secret185[i] = 0
	}
	backtrack185(0)
}

func solve() int64 {
	once185.Do(compute185)
	return answer185
}

func main() { bench.Run(185, solve) }
