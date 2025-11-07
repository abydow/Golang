package main

import "fmt"

func main() {
	var x int = 32
	var y float64 = float64(x)
	z := uint(y)
	fmt.Printf("%T and %T and %T\n", x, y, z)
}
