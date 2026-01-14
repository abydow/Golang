package main

import "fmt"

func main() {
	fmt.Println("Exercise 4")

	xf := []int{42,43,44,45,46,47,48,49,50,51}
	fmt.Println(xf)
	xf = append(xf, 52)
	fmt.Println(xf)
	xf = append(xf, 53, 54, 55)
	fmt.Println(xf)

	y:= []int{56, 57, 58, 59, 60}
	xf = append(xf, y...)
	fmt.Println(xf)

}
