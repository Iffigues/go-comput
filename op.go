package main

import "fmt"

func (r *obj) getPoly(a int) (b float64) {
	for _, val := range r.finalPoly {
		if val.degree == a {
			return val.nb
		}
	}
	return 0
}

func (r *obj) zero() {
	r.print()
	r.getPoly(0)
}

func (r *obj) one() {
	r.print()
	a, b := ireductible(r.getPoly(0)*-1, r.getPoly(1))
	if a != r.getPoly(0)*-1 || b != r.getPoly(1) {
		fmt.Println("Fraction irreductibl: ", a, " / ", b)
	}
	fmt.Println(a / b)
}

func (r *obj) twos(a, b float64) {
	b = b * -1
	aa, bb := ireductible(a, b)
	if aa != a || bb != b {
		a = aa
		b = bb
	}
	dis, err := Sqrt(b / a)
	if err != nil && err.Error() == "it's negatice" {
		fmt.Println(a, b, "ezez")
	} else {
		fmt.Println(dis, "ez")
	}
}

func (r *obj) two() {
	r.print()
	c := r.getPoly(0)
	b := r.getPoly(1)
	a := r.getPoly(2)
	if b == 0 {
		r.twos(a, c)
		return
	}
	discriminant := disc(b, a, c)
	if discriminant > 0 {
		dis, _ := Sqrt(discriminant)
		x1 := -b + dis
		x2 := -b - dis
		y := 2 * a
		g1, g2 := ireductible(x1, y)
		g3, g4 := ireductible(x2, y)
		fmt.Println(discriminant)
		if g1 != x1 || g2 != y {
			fmt.Println("fi", g1, "/", g2)
			x1 = g1 / g2
		} else {
			x1 = x1 / y
		}
		if g3 != x2 || g4 != y {
			fmt.Println(x2, "fa", g3, "/", g4)
			x2 = g3 / g4
		} else {
			x2 = x2 / y
		}
		fmt.Println(x1, x2)

	} else if discriminant == 0 {

	} else {

	}
}

func (r *obj) other() {
	r.print()
}
