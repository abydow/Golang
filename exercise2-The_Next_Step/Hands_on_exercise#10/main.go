package main

import "fmt"

var y int

func main() {
	fmt.Printf("- - - %d - - -  \n", y) //for 0 value

	z := 10.00

	fmt.Printf("%f \t %T \n", z, z) //short declaration

	a, b := 10, 20
	fmt.Println(a, b) //multi=assignment

	var p complex128 = -1 * .5
	fmt.Println(p)

	const (
		_ = iota
		w
		x
		y
	)

	fmt.Println(w, x, y)

}
