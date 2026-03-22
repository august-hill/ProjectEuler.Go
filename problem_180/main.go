// Problem 180: Rational Zeros of a Function of Three Variables (Golden Triplets)
// Answer: 285196020571078987

package main

import (
	"math"
	"sync"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

const K180 = 35

type rat180 struct{ n, d int64 }

func gcd180(a, b int64) int64 {
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

func mkrat180(n, d int64) rat180 {
	if d < 0 {
		n, d = -n, -d
	}
	if n == 0 {
		return rat180{0, 1}
	}
	g := gcd180(abs180(n), d)
	return rat180{n / g, d / g}
}

func abs180(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func radd180(a, b rat180) rat180 {
	return mkrat180(a.n*b.d+b.n*a.d, a.d*b.d)
}

func rvalid180(r rat180) bool {
	return r.n > 0 && r.d > 0 && r.n < r.d && r.d <= K180
}

const hsz180 = 1 << 20
const hmask180 = hsz180 - 1

type htEntry180 struct {
	val  rat180
	used bool
}

var ht180 [hsz180]htEntry180

func hclear180() {
	for i := range ht180 {
		ht180[i].used = false
	}
}

func hins180(r rat180) bool {
	h := uint64(r.n*100003+r.d*999983) & hmask180
	for ht180[h].used {
		if ht180[h].val == r {
			return false
		}
		h = (h + 1) & hmask180
	}
	ht180[h].val = r
	ht180[h].used = true
	return true
}

var (
	once180        sync.Once
	answerCache180 int64
)

func compute180() {
	hclear180()

	rats := make([]rat180, 0, 700)
	for q := int64(2); q <= K180; q++ {
		for p := int64(1); p < q; p++ {
			if gcd180(p, q) == 1 {
				rats = append(rats, mkrat180(p, q))
			}
		}
	}

	total := rat180{0, 1}

	for i := 0; i < len(rats); i++ {
		for j := i; j < len(rats); j++ {
			x := rats[i]
			y := rats[j]

			// n=1: z = x+y
			{
				z := radd180(x, y)
				if rvalid180(z) {
					s := radd180(radd180(x, y), z)
					if hins180(s) {
						total = radd180(total, s)
					}
				}
			}

			// n=-1: z = xy/(x+y)
			{
				zn := x.n * y.n
				zd := x.n*y.d + y.n*x.d
				if zd > 0 && zn > 0 {
					z := mkrat180(zn, zd)
					if rvalid180(z) {
						s := radd180(radd180(x, y), z)
						if hins180(s) {
							total = radd180(total, s)
						}
					}
				}
			}

			// n=2: z = sqrt(x^2+y^2)
			{
				num2 := x.n*x.n*y.d*y.d + y.n*y.n*x.d*x.d
				den1 := x.d * y.d
				sq := int64(math.Sqrtf(float32(num2)))
				for sq*sq < num2 {
					sq++
				}
				for sq*sq > num2 {
					sq--
				}
				if sq > 0 && sq*sq == num2 {
					z := mkrat180(sq, den1)
					if rvalid180(z) {
						s := radd180(radd180(x, y), z)
						if hins180(s) {
							total = radd180(total, s)
						}
					}
				}
			}

			// n=-2: z = xy/sqrt(x^2+y^2)
			{
				num2 := x.n*x.n*y.d*y.d + y.n*y.n*x.d*x.d
				sq := int64(math.Sqrtf(float32(num2)))
				for sq*sq < num2 {
					sq++
				}
				for sq*sq > num2 {
					sq--
				}
				if sq > 0 && sq*sq == num2 {
					z := mkrat180(x.n*y.n, sq)
					if rvalid180(z) {
						s := radd180(radd180(x, y), z)
						if hins180(s) {
							total = radd180(total, s)
						}
					}
				}
			}

			if i != j {
				for ord := 0; ord < 2; ord++ {
					var big, small rat180
					if ord == 0 {
						big, small = y, x
					} else {
						big, small = x, y
					}
					if big.n*small.d <= small.n*big.d {
						continue
					}

					// n=1 diff: z = big - small
					{
						z := mkrat180(big.n*small.d-small.n*big.d, big.d*small.d)
						if rvalid180(z) {
							s := radd180(radd180(big, small), z)
							if hins180(s) {
								total = radd180(total, s)
							}
						}
					}

					// n=2 diff: z = sqrt(big^2 - small^2)
					{
						num2 := big.n*big.n*small.d*small.d - small.n*small.n*big.d*big.d
						if num2 > 0 {
							den1 := big.d * small.d
							sq := int64(math.Sqrtf(float32(num2)))
							for sq*sq < num2 {
								sq++
							}
							for sq*sq > num2 {
								sq--
							}
							if sq > 0 && sq*sq == num2 {
								z := mkrat180(sq, den1)
								if rvalid180(z) {
									s := radd180(radd180(big, small), z)
									if hins180(s) {
										total = radd180(total, s)
									}
								}
							}
						}
					}

					// n=-1 diff: z = big*small/(big-small)
					{
						zn := big.n * small.n
						zd := big.n*small.d - small.n*big.d
						if zd > 0 && zn > 0 {
							z := mkrat180(zn, zd)
							if rvalid180(z) {
								s := radd180(radd180(big, small), z)
								if hins180(s) {
									total = radd180(total, s)
								}
							}
						}
					}

					// n=-2 diff: z = big*small/sqrt(big^2 - small^2)
					{
						num2 := big.n*big.n*small.d*small.d - small.n*small.n*big.d*big.d
						if num2 > 0 {
							sq := int64(math.Sqrtf(float32(num2)))
							for sq*sq < num2 {
								sq++
							}
							for sq*sq > num2 {
								sq--
							}
							if sq > 0 && sq*sq == num2 {
								z := mkrat180(big.n*small.n, sq)
								if rvalid180(z) {
									s := radd180(radd180(big, small), z)
									if hins180(s) {
										total = radd180(total, s)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	answerCache180 = total.n + total.d
}

func solve() int64 {
	once180.Do(compute180)
	return answerCache180
}

func main() { bench.Run(180, solve) }
