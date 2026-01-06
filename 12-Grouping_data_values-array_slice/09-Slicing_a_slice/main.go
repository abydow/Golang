package main

import (
	"fmt"
)

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(xi[ : ])
	fmt.Println("-------------------")
	// [inclusive : exclusive]
	fmt.Println(xi[2:6])
	fmt.Println("-------------------")
	// [:exclusive]
	fmt.Println(xi[:6])
	fmt.Println("-------------------")
	// [inclusive:]
	fmt.Println(xi[2:])
	fmt.Println("-------------------")
}
