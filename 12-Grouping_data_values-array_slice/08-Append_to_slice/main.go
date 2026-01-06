package main

import (
	"fmt"
)

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// variadic parameter
	xi = append(xi, 52, 53, 54)
	fmt.Println(xi)
}
