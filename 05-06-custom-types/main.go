package main

import "fmt"

// Custom types
type manzana int
type pera int

func main() {
	var m manzana
	var p pera
	m = 10
	p = 20
	a := m + manzana(p)
	b := pera(m) + p
	fmt.Println(a, b)
	fmt.Printf("%T %T", a, b)
}
