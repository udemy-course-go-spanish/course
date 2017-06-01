package main

import "fmt"

func main() {
	// string
	// secuencia de bytes
	var a string
	a = "Hola Gophers!"
	b := make([]byte, 4)
	b[0] = 72
	b[1] = 111
	b[2] = 108
	b[3] = 65
	fmt.Printf("%T, %v, %#v\n", a, a, a)
	fmt.Println([]byte(a))
	fmt.Println(string(b))
}
