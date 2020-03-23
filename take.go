package main

func (r *obj) getDegree() (a int) {
	a = 0
	for _, val := range r.finalPoly {
		if val.degree > a {
			a = val.degree
		}
	}
	return
}

func (r *obj) calc() {
	f := r.getDegree()
	if f == 0 {
		r.zero()
	} else if f == 1 {
		r.one()
	} else if f == 2 {
		r.two()
	} else {

	}
}
