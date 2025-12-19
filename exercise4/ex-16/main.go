package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i <= 100; i++ {

		if x:= rand.Intn(5); x == 3{
			fmt.Printf("x is equal to 3: %d\n", x)
		}
	}
}
