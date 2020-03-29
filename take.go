package main

import "fmt"

func (r *obj) getDegree() (a int) {
	a = 0
	for _, val := range r.finalPoly {
		if val.degree > a {
			a = val.degree
		}
	}
	return
}

func (r *obj) is_possible() (ok bool) {
	for _, val := range r.finalPoly {
		if val.degree < 0 {
			return false
		}
	}
	return true
}

func (r *obj) calc() {
	f := r.getDegree()
	if !r.is_possible() {

	}
	if len(r.finalPoly) == 0 {
		fmt.Println("all solution")
	} else if f == 0 {
		r.zero()
	} else if f == 1 {
		r.one()
	} else if f == 2 {
		r.two()
	} else {
		r.other()
	}
}
