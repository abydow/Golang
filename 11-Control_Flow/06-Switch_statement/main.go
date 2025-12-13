package main

import (
	"fmt"
)

func main(){
	fmt.Println("Enter a number between 1 to 3:")
	var x int
	fmt.Scanln(&x)
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other than one, two, three")
	}

	// to check every condition
	y:= 2
	switch y {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two")
		fallthrough
	case 3:				// will execute even if case 3 is not matched
		fmt.Println("three")
		fallthrough
	default: 			// will execute even if default is not matched
		fmt.Println("other than one, two, three")
	}
}
