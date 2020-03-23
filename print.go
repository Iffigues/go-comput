package main

import "fmt"

func (r *obj) print() {
	for key, val := range r.finalPoly {
		h := val.nb
		if key > 0 {
			if val.nb > 0 {
				print("+ ")
			} else {
				print("- ")
				h = h * -1
			}
		}
		fmt.Print(h)
		if val.degree > 0 {
			print(" * X")
		}
		if val.degree > 1 {
			print(" ^ ")
			print(val.degree)
		}
		print(" ")
	}
	print("= 0\n")
}
