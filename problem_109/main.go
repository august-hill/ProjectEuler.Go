// Problem 109: Darts
// How many distinct ways can a player checkout with a score less than 100?
// Answer: 38182

package main

import "github.com/august-hill/ProjectEuler.Go/bench"

func solve() int64 {
	// Singles: 1-20, 25, Doubles: 2-40, 50, Trebles: 3-60
	var singles [62]int
	ndarts := 0
	for i := 1; i <= 20; i++ {
		singles[ndarts] = i
		ndarts++
	}
	singles[ndarts] = 25
	ndarts++
	for i := 1; i <= 20; i++ {
		singles[ndarts] = 2 * i
		ndarts++
	}
	singles[ndarts] = 50
	ndarts++
	for i := 1; i <= 20; i++ {
		singles[ndarts] = 3 * i
		ndarts++
	}

	var doubles [21]int
	ndoubles := 0
	for i := 1; i <= 20; i++ {
		doubles[ndoubles] = 2 * i
		ndoubles++
	}
	doubles[ndoubles] = 50
	ndoubles++

	count := 0

	// 1 dart: just a double
	for i := 0; i < ndoubles; i++ {
		if doubles[i] < 100 {
			count++
		}
	}

	// 2 darts: any + double
	for i := 0; i < ndarts; i++ {
		for j := 0; j < ndoubles; j++ {
			if singles[i]+doubles[j] < 100 {
				count++
			}
		}
	}

	// 3 darts: any + any (ordered pair with i<=j) + double
	for i := 0; i < ndarts; i++ {
		for j := i; j < ndarts; j++ {
			for k := 0; k < ndoubles; k++ {
				if singles[i]+singles[j]+doubles[k] < 100 {
					count++
				}
			}
		}
	}

	return int64(count)
}

func main() { bench.Run(109, solve) }
