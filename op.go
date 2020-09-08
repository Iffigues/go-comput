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
		fmt.Println("Fraction irreductible: ", a, " / ", b)
	}
	fmt.Println(a / b)
}

func (r *obj) twos(a, b float64) {
	b = b * -1
	aa, bb := ireductible(a, b)
	if (aa != a || bb != b) {
		a = aa
		b = bb
		fmt.Println("Fraction irreductible: ", a, " / ", b)
	}
	dis, err := Sqrt(b / a)
	if err != nil && err.Error() == "it's negatice" {
		fmt.Println("the solution was  √(", a, "/", b, ")")
	} else {
		fmt.Println("the solution was ", dis)
	}
}

func (r *obj) two() {
	r.print()
	c := r.getPoly(0)
	b := r.getPoly(1)
	a := r.getPoly(2)
	fmt.Println(c,b,a)
	if b == 0 {
		r.twos(a, c)
		return
	}
	dis := disc(b, a, c)
	if dis > 0 {
		dic, _  := Sqrt(dis)
		x1 := -b + dic
		x2 := -b - dic
		fmt.Println(x1,x2)
		y := 2 * a
		g1, g2 := ireductible(x1, y)
		g3, g4 := ireductible(x2, y)
		if g1 != x1 || g2 != y  {
			fmt.Println("Fraction irreductible: ", g1, "/", g2)
			x1 = g1 / g2
		} else {
			fmt.Println("hallo")
			x1 = x1 / y
		}
		if g3 != x2 || g4 != y {
			fmt.Println(x2, "Fraction irreductible: ", g3, "/", g4)
			x2 = g3 / g4
		} else {
			fmt.Println("hello")
			x2 = x2 / y
		}
		fmt.Println("the first solution was ", x1)
		fmt.Println("the two solution was ", x2)

	} else if dis == 0 {
		 y := 2 * a
		r1, r2 := ireductible(b, y)
		if b != r1 || r2 != y {
			fmt.Println("Fraction irreductible: ", r1, "/", r2)
			r1 = -(r1 / r2)
		} else {
			r1 = -(b/y)
		}
		fmt.Println("the solutoin was ", r1)
	} else {
		fmt.Println("not solution")
	}
}

func (r *obj) other() {
	r.print()
}
