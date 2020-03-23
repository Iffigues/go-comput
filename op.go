package main

func (r *obj) getPoly(a int) (b float64) {
	for _, val := range r.finalPoly {
		if val.degree == a {
			return val.nb
		}
	}
	return 0
}

func (r *obj) zero() {
	r.getPoly(0)
}

func (r *obj) one() {
	r.getPoly(0)
	r.getPoly(1)
}

func (r *obj) two() {
	r.getPoly(0)
	r.getPoly(1)
	r.getPoly(2)
}
