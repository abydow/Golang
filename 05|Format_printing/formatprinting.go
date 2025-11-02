package main

import (
	"fmt"
)

func main() {
	const name = "Kim"
	const age = 22

	// const name, age = "kim", 22

	fmt.Printf("%s is %d years old. \n", name, age)
	//fmt.Printf("%s is %d years old. \t and the types are %T and%T", name, age, name, age)
	fmt.Printf("%v is %v years old. \t and the types are %T and %T ", name, age, name, age)
}
