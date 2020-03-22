package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func (c *obj) addSpace(a string) (b string) {
	a = strings.ReplaceAll(a, "x", "X")
	for _, zz := range a {
		z := string(zz)
		r := strings.ContainsAny(z, "+*-/^=X")
		if r {
			b = b + " "
		}
		b = b + z
		if r {
			b = b + " "
		}
	}
	return
}

func (r *obj) getString(a string) (j string) {
	c := strings.ReplaceAll(a, "+ +", "+")
	d := strings.ReplaceAll(c, "+ -", "-")
	e := strings.ReplaceAll(d, "- -", "+")
	f := strings.ReplaceAll(e, "- +", "-")
	g := strings.ReplaceAll(f, " X", "X")
	gg := strings.ReplaceAll(g, " *", "")
	h := strings.ReplaceAll(gg, " ^ ", "")
	i := strings.ReplaceAll(h, "+ ", "+")
	j = strings.ReplaceAll(i, "- ", "-")
	return j
}

func (r *obj) buildPoly(a string) (p poly) {
	f := strings.Split(a, "X")
	if len(f) > 2 {
		log.Fatal("error in process")
	}
	if len(f) == 1 {
		f = append(f, "0")
	}
	if f[0] == "" {
		f[0] = "1"
	}
	if f[1] == "" {
		f[1] = "1"
	}
	var err error
	if p.nb, err = strconv.ParseFloat(f[0], 64); err != nil {
		log.Fatal("error in conversion")
	}
	if p.degree, err = strconv.Atoi(f[1]); err != nil {
		log.Fatal("error in conversion")
	}
	return
}

func (r *obj) makePoly(a string) (b []poly) {
	j := strings.Split(r.getString(a), " ")
	for _, val := range j {
		b = append(b, r.buildPoly(val))
	}
	fmt.Println(b)
	return
}

func (r *obj) verify(a string) (b []string) {
	a = strings.Join(strings.Fields(r.addSpace(a)), " ")
	return strings.Split(a, "=")
}

func (r *obj) makeInt() {
	for _, val := range r.poly1 {
		r.intPoly = append(r.intPoly, val)
	}
	for _, val := range r.poly2 {
		val.nb = val.nb * -1
		r.intPoly = append(r.intPoly, val)
	}
}

func (r *obj) makeMiddle() {
	for _, val := range r.intPoly {
		if _, ok := r.middlePoly[val.degree]; !ok {
			r.middlePoly[val.degree] = nil
		}
	}
	for val, _ := range r.middlePoly {
		for _, value := range r.intPoly {
			if value.degree == val {
				r.middlePoly[val] = append(r.middlePoly[val], value.nb)
			}
		}
	}
}

func (r *obj) makeFinal() {
	for key, val := range r.middlePoly {
		var t float64
		t = 0
		for _, value := range val {
			t = t + value
		}
		if t != 0 {
			r.finalPoly = append(r.finalPoly, poly{t, key})
		}
	}
	sort.Sort(ByAge(r.finalPoly))
	fmt.Println(r.finalPoly)
}

func newObj(a []string) (r *obj, err error) {

	if a == nil {
		return nil, errors.New("empty string")
	}

	r = new(obj)
	r.middlePoly = make(map[int][]float64)
	ops := r.verify(strings.Join(a, " "))
	if len(ops) != 2 {
		return nil, errors.New("error in equation")
	}
	r.poly1 = r.makePoly(strings.Trim(ops[0], " "))
	r.poly2 = r.makePoly(strings.Trim(ops[1], " "))
	r.makeInt()
	r.makeMiddle()
	r.makeFinal()
	return
}
