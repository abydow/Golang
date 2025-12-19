package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 15; i++ {
		x := rand.Intn(10)
		fmt.Printf("x: %d -> \n", x)
		if x == 0 { // loops until 0 shows up
			break
		}
	}
}

