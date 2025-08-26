package main

import "fmt"

func main() {
	const (
		a = iota
		b = 2
		c = 3
		d
		e
		f
		g
	)
	fmt.Println("a:", a, "b:", b, "c:", c, "d:", d, "e:", e, "f:", f, "g:", g)
	const (
		h = iota
		i
		j
	)
	fmt.Println("h:", h, "i:", i, "j:", j)
}
