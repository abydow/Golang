package main

import "fmt"

func main() {
	v := 42 - 5i //variable type is inferred  from the value on right hand side
	fmt.Printf("v is of type %T\n", v)

}
