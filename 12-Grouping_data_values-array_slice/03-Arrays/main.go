package main

import (
	"fmt"
)

func main(){
	numbers := [10]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(numbers)

	names := [...]string{"Alice", "Bob", "Charlie", "Diana"}
	fmt.Println(names)

	var c [2]int
	c[0]=100
	c[1]=200
	fmt.Printf("%T\n", c)
}
