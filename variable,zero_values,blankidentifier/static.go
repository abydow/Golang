package main

import "fmt"

func main() {

	a := 12
	fmt.Println(a)

	a, b, _, c := 1, 2, 3, "home"
	fmt.Println(a, b, c)

	//thing that will not work
	/*
		a, b, c, d = 1, 2, 3, 4
		fmt.Println(a,b,c)
	*/

	var g int
	fmt.Println(g)
	g = 42
	fmt.Println(g)

	//Declare a variable to hold the value of a certain type
	//Then assign a value to that variable
	//initialize a variable

	var h int = 100
	fmt.Println(h)
}
