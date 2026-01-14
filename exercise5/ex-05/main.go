package main

import "fmt"

func main() {
	fmt.Println("Exercise 5")
	xf := []int{42,43,44,45,46,47,48,49,50,51}

	xf = append(xf[ :3], xf[6: ]...)
	fmt.Println(xf)}
