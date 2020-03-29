package main

func square(b, h float64) (a float64) {
	a = b
	for ; h > 1; h-- {
		a = a * b
	}
	return
}

func disc(b, a, c float64) (dix float64) {
	return square(b, 2.0) - (4 * a * c)
}
