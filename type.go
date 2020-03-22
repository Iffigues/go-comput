package main

type poly struct {
	nb     float64
	degree int
}

type ByAge []poly

type obj struct {
	poly1      []poly
	poly2      []poly
	intPoly    []poly
	middlePoly map[int][]float64
	finalPoly  []poly
}
