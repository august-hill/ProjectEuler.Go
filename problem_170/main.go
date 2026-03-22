// Problem 170: Find the largest 0 to 9 pandigital concatenated product
// Answer: 9857164023

package main

import (
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

var (
	once170        sync.Once
	answerCache170 int64
)

func isPandigital170(num int64) bool {
	if num < 1000000000 || num > 9999999999 {
		return false
	}
	var seen [10]bool
	for num > 0 {
		d := num % 10
		if seen[d] {
			return false
		}
		seen[d] = true
		num /= 10
	}
	for i := 0; i < 10; i++ {
		if !seen[i] {
			return false
		}
	}
	return true
}

func concatLL170(a, b int64) int64 {
	tmp := b
	if tmp == 0 {
		return a*10 + b
	}
	for tmp > 0 {
		a *= 10
		tmp /= 10
	}
	return a + b
}

// checkFactor: try all permutations of digsF, check if n*f has digits digsNF
func checkFactor170(n int64, digsF []int, digsNF []int) int64 {
	arr := make([]int, len(digsF))
	copy(arr, digsF)
	// Sort arr
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	nf := len(arr)
	nnf := len(digsNF)

	for {
		if nf <= 1 || arr[0] != 0 {
			f := int64(0)
			for _, d := range arr {
				f = f*10 + int64(d)
			}
			nfVal := n * f
			var check [10]int
			tmp := nfVal
			cntD := 0
			for tmp > 0 {
				check[tmp%10]++
				tmp /= 10
				cntD++
			}
			if cntD == nnf {
				var sortedNF [10]int
				si := 0
				for d := 0; d <= 9; d++ {
					for c := 0; c < check[d]; c++ {
						sortedNF[si] = d
						si++
					}
				}
				ok := true
				for i := 0; i < nnf; i++ {
					if sortedNF[i] != digsNF[i] {
						ok = false
						break
					}
				}
				if ok {
					return f
				}
			}
		}

		// Next permutation
		j := nf - 2
		for j >= 0 && arr[j] >= arr[j+1] {
			j--
		}
		if j < 0 {
			break
		}
		l := nf - 1
		for arr[l] <= arr[j] {
			l--
		}
		arr[j], arr[l] = arr[l], arr[j]
		lo, hi := j+1, nf-1
		for lo < hi {
			arr[lo], arr[hi] = arr[hi], arr[lo]
			lo++
			hi--
		}
	}
	return -1
}

func compute170() {
	best := int64(0)

	for n := 2; n <= 98; n++ {
		// Check n has no repeated digits
		var nDigs [3]int
		nn := 0
		tmp := n
		for tmp > 0 {
			nDigs[nn] = tmp % 10
			nn++
			tmp /= 10
		}
		var nSeen [10]bool
		validN := true
		for i := 0; i < nn; i++ {
			if nSeen[nDigs[i]] {
				validN = false
				break
			}
			nSeen[nDigs[i]] = true
		}
		if !validN {
			continue
		}

		// Remaining digits for factors
		var remDigs [10]int
		nrem := 0
		for d := 0; d <= 9; d++ {
			if !nSeen[d] {
				remDigs[nrem] = d
				nrem++
			}
		}

		nsub := 1 << nrem
		for mask1 := 1; mask1 < nsub-1; mask1++ {
			var digsF1 [10]int
			var digsF2 [10]int
			nf1, nf2 := 0, 0
			for i := 0; i < nrem; i++ {
				if mask1&(1<<i) != 0 {
					digsF1[nf1] = remDigs[i]
					nf1++
				} else {
					digsF2[nf2] = remDigs[i]
					nf2++
				}
			}

			maxE := 10 - (nf1 + nf2) // = nn
			for e1 := 0; e1 <= maxE; e1++ {
				nf1p := nf1 + e1
				nf2p := 10 - nf1p
				if nf2p < nf2 || nf2p > nf2+maxE {
					continue
				}
				if nf1p <= 0 || nf2p <= 0 {
					continue
				}

				// Enumerate combinations of size nf1p from {0..9}
				nc := nf1p
				if nc > 10 || 10-nc < nf2p {
					continue
				}
				combo := make([]int, nc)
				for i := 0; i < nc; i++ {
					combo[i] = i
				}
				for {
					var inCombo [10]bool
					for i := 0; i < nc; i++ {
						inCombo[combo[i]] = true
					}
					var digsNF1 [10]int
					var digsNF2 [10]int
					n1, n2 := 0, 0
					for d := 0; d <= 9; d++ {
						if inCombo[d] {
							digsNF1[n1] = d
							n1++
						} else {
							digsNF2[n2] = d
							n2++
						}
					}

					f1 := checkFactor170(int64(n), digsF1[:nf1], digsNF1[:n1])
					if f1 >= 0 {
						f2 := checkFactor170(int64(n), digsF2[:nf2], digsNF2[:n2])
						if f2 >= 0 {
							p1 := int64(n) * f1
							p2 := int64(n) * f2
							prodConcat := concatLL170(p1, p2)
							if isPandigital170(prodConcat) && prodConcat > best {
								best = prodConcat
							}
						}
					}

					// Next combination
					i := nc - 1
					for i >= 0 && combo[i] == 10-nc+i {
						i--
					}
					if i < 0 {
						break
					}
					combo[i]++
					for j := i + 1; j < nc; j++ {
						combo[j] = combo[j-1] + 1
					}
				}
			}
		}
	}
	answerCache170 = best
}

func solve() int64 {
	once170.Do(compute170)
	return answerCache170
}

func main() { bench.Run(170, solve) }
