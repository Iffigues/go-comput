package main

func square(b, h float64) (a float64) {
	a = b
	for ; h > 1; h-- {
		a = a * b
	}
	return
}

func disc(b, a, c float64) (dix float64) {
	ee := square(b, 2.0)
	aa := 4.0 * a * c
	return ee - aa
}

func null(b, a float64) (dix float64) {
	i := b / (a)
	return -i
}
