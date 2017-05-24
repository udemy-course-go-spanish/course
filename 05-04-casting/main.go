package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		var a int
		var b float32
		a = 15
		b = 32.9999
		c := float32(a) + b
		fmt.Println(c)
		fmt.Printf("%T\n", c)
	*/

	/*
		var a string
		var b int
		a = "74"
		b = 1
		d, _ := strconv.Atoi(a)
		c := d + b
		fmt.Println(c)
	*/

	var a string
	var b int
	a = "Cantidad de usuarios "
	b = 50
	c := a + strconv.Itoa(b)
	fmt.Println(c)
}
