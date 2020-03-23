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
	r.print()
	r.getPoly(0)
}

func (r *obj) one() {
	r.print()
	r.getPoly(0)
	r.getPoly(1)
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
