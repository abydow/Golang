package main

import (
	"fmt"
)

func main() {
	fmt.Println("Exercise 2")

	xf := []int{42,43,44,45,46,47,48,49,50,51}

	for i,v := range xf {
		fmt.Printf("%v\t%T\t%v\n",v,v,i)
	}
}
