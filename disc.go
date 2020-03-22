package main

func square(b, h int) (a int) {
	a = b
	for ; h > 1; h-- {
		a = a * b
	}
	return
}

func disc(b, a, c int) (dix int) {
	return square(b, 2) - (4 * a * c)
}
