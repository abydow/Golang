package main

import (
	"fmt"
)

func main() {
	fmt.Println("Deleting from a slice")
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(xi)

	// Deleting the value at index 3 (value 45)
	xf := append(xi[:3], xi[4:]...)
	fmt.Println(xf)
}
