package main

import "fmt"

func main() {
	// zero value
	var a int
	var b float32
	var c bool
	var d string
	fmt.Println(a, b, c, d)
	fmt.Printf("%q\n", d)
}
