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
	if a != r.getPoly(0)*-1 && b != r.getPoly(1) {
		fmt.Println("Fraction irreductibl: ", a, " / ", b)
	}
	fmt.Println(a / b)
}

func (r *obj) two() {
	r.print()
	r.getPoly(0)
	r.getPoly(1)
	r.getPoly(2)
}

func (r *obj) other() {
	r.print()
}
