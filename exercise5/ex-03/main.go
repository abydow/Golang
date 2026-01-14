package main

import (
	"fmt"
)

func main() {
	fmt.Println("Exercise 3")

	xf := []int{42,43,44,45,46,47,48,49,50,51}

	w := xf[0:5]
	x := xf[5:10]
	y := xf[2:7]
	z := xf[1:6]

	fmt.Println("w:", w)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
