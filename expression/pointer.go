package main

func main() {
	x := 10
	var p *int = &x
	*p += 20
	println(p, *p)

	p2 := &x
	println(p == p2)

	s := struct {
		x int
	}{}
	s.x = 100
	ps := s
	ps.x += 100
	println(ps.x)
}
