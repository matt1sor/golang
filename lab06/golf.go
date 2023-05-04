package main

import (
	q "flag"
	r "fmt"
	u "math"
)

func main() {
	p := r.Print
	l := q.Float64
	a := l("a", 1, "")
	b := l("b", 1, "")
	c := l("c", 1, "")
	q.Parse()
	x, y, z := *a, *b, *c
	d := y*y - 4*(x*z)

	if d >= 0 {
		m := (-y + u.Sqrt(d)) / (2 * x)
		p(m, m)
	} else {
		p("F")
	}
}
