package main

import (
	"math"
)

func root(a float64, n int) float64 {
	n1 := n - 1
	n1f, rn := float64(n1), 1/float64(n)
	x, x0 := 1., 0.
	for {
		potx, t2 := 1/x, a
		for b := n1; b > 0; b >>= 1 {
			if b&1 == 1 {
				t2 *= potx
			}
			potx *= potx
		}
		x0, x = x, rn*(n1f*x+t2)
		if math.Abs(x-x0)*1e15 < x {
			break
		}
	}
	return x
}

func equalByTolerance(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= a*tolerance
}

func findNearestBase10(v float64) float64 {
	goUp := v < 1
	divs := 0
	for {
		if v <= 10 && v >= 1 {
			break
		}
		if goUp {
			v *= 10
		} else {
			v /= 10
		}
		divs++
	}
	num := math.Pow10(divs)
	if goUp {
		return 1 / num
	}
	return num
}
