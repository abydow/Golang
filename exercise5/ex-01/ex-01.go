package main

import (
	"fmt"
)

func main() {
	fmt.Println("Exercise 1")


	xf := [5]int{1,2,3,4,5}

	for i,x := range xf {
		fmt.Printf("%v\t%T\t%v\n",x,x,i)
	}
}
