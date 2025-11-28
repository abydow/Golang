package main

import "fmt"

var x int = 42

const y float32 = 29.830

func main() {
	z := "Aby"
	fmt.Printf("%v is %T type\n", x, x)
	fmt.Printf("%v is %T type\n", y, y)
	fmt.Printf("%v is %T type\n", z, z)
}
