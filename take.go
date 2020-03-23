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

	} else if f == 1 {

	} else if f == 2 {

	} else {

	}
}
