package main

func (r *obj) print() {
	for key, val := range r.finalPoly {
		h := val.nb
		if key > 0 {
			if val.nb > 0 {
				print(" + ")
			} else {
				print(" - ")
			}
			h = h * -1
		}
		print(h)
		if val.degree > 0 {
			print("X")
		}
		if val.degree > 0 {
			print(" ^ ")
			print(val.degree)
		}
		print(" ")
	}
	print("\n")
}
