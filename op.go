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

func (r *obj) two() {
	r.print()
	a := r.getPoly(0)
	b := r.getPoly(1)
	c := r.getPoly(2)
	discriminant := disc(b, a, c)
	if discriminant > 0 {

	} else if discriminant == 0 {

	} else {

	}
}

func (r *obj) other() {
	r.print()
}
