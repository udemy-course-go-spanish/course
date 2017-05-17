package main

import "fmt"

func main() {
	a, _, _ := devolver()
	fmt.Println(a)
}

func devolver() (int, int, int) {
	return 2, 9, 6
}
