package main

import "fmt"

var g int = 42

func Newfunc() {
	fmt.Println(g)
	y := 42.02
	fmt.Println(y)

}

func Nextfunc() {
	// fmt.Println(g) // will print 42 cause g is a global variable
	// fmt.Println (y) //error cause y is defined in previous function not this function (local scope)
}

func main() {
	Newfunc()
	Nextfunc()
}
