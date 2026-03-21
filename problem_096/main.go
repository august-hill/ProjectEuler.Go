// Problem 096: Su Doku
// Solve all 50 Sudoku puzzles and sum the 3-digit numbers in the top-left corners.
// Answer: 24702

package main

import (
	_ "embed"
	"strings"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

//go:embed p096_sudoku.txt
var sudokuData096 string

type board096 [9][9]int

func getPossible096(b *board096, r, c int) int {
	used := 0
	for j := 0; j < 9; j++ {
		used |= 1 << uint(b[r][j])
	}
	for i := 0; i < 9; i++ {
		used |= 1 << uint(b[i][c])
	}
	br, bc := (r/3)*3, (c/3)*3
	for i := br; i < br+3; i++ {
		for j := bc; j < bc+3; j++ {
			used |= 1 << uint(b[i][j])
		}
	}
	return ^used & 0x3FE
}

func popcount096(x int) int {
	count := 0
	for x != 0 {
		count++
		x &= x - 1
	}
	return count
}

func solveSudoku096(b *board096) bool {
	minOpts := 10
	bestR, bestC := -1, -1
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if b[r][c] == 0 {
				opts := popcount096(getPossible096(b, r, c))
				if opts < minOpts {
					minOpts = opts
					bestR, bestC = r, c
					if opts == 1 {
						goto found
					}
				}
			}
		}
	}
found:
	if bestR == -1 {
		return true // solved
	}
	if minOpts == 0 {
		return false // dead end
	}

	possible := getPossible096(b, bestR, bestC)
	for d := 1; d <= 9; d++ {
		if possible&(1<<uint(d)) != 0 {
			b[bestR][bestC] = d
			if solveSudoku096(b) {
				return true
			}
			b[bestR][bestC] = 0
		}
	}
	return false
}

func solve() int64 {
	lines := strings.Split(strings.TrimSpace(sudokuData096), "\n")
	total := 0

	for puzzle := 0; puzzle < 50; puzzle++ {
		base := puzzle*10 + 1
		var b board096
		for r := 0; r < 9; r++ {
			line := strings.TrimSpace(lines[base+r])
			for c := 0; c < 9; c++ {
				b[r][c] = int(line[c] - '0')
			}
		}
		solveSudoku096(&b)
		total += b[0][0]*100 + b[0][1]*10 + b[0][2]
	}
	return int64(total)
}

func main() { bench.Run(96, solve) }
