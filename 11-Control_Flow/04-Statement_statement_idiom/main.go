package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x:=20
	y:=40

	if z:=2*rand.Intn(x); z<y { // writing the statement inside the if condition
		fmt.Printf(" %v is less than %v\n", z, y)
	} else{
		fmt.Printf(" %v is equal to %v\n", z, y)
	}
}
