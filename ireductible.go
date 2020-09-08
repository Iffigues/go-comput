package main

func is_whole(a float64) (ok bool) {
	if a == float64(int64(a)) {
		return true
	}
	return false
}

func find(a, b float64) (c, d float64, ok bool) {
	g, f := a, b
	for f > 1 {
		if is_whole(b/f) && is_whole(g/f) {
			return g / f, b / f, true
		}
		f = f - 1
	}
	return a, b, false
}

func ireductible(a, b float64) (c, d float64) {
	sa, sb := 1.0, 1.0
	if a == 0 {
		return a, b
	}
	if a < 0 {
		sa = -1.0
		a = a * -1.0
	}
	if b < 0 {
		sb = -1.0
		b = b * -1.0
	}
	if !is_whole(a) || !is_whole(b) {
		return a * sa, b * sb
	}
	if is_whole(a / b) {
		return a * sa / b * sb, b * sb / b * sb
	}
	ok := true
	for {
		if !ok {
			return sa * a, sb * b
		}
		a, b, ok = find(a, b)
	}
}
