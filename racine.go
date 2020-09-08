package main

import "errors"

func abs(a float64) float64 {
	if a < 0 {
		return -1 * a
	}
	return a
}

func Sqrt(x float64) (z float64, err error) {
	if x == 0 {
		return x, errors.New("can't be divide by 0")
	}
	if x < 0 {
		return x, errors.New("it's negatice")
	}
	z = 100.0
	delta := 0.0000001
	step := func() float64 {
		return z - (z*z-x)/(2*z)
	}
	for zz := step(); abs(zz-z) > delta; {
		z = zz
		zz = step()
	}
	return
}
