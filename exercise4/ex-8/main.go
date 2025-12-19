package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 42; i++ {
		x := rand.Intn(5)

		switch x {
		case 0:
			fmt.Printf("%d: Zero | Iteration no. %v\n", x, i)
		case 1:
			fmt.Printf("%d: One | Iteration no. %v\n", x, i)
		case 2:
			fmt.Printf("%d: Two | Iteration no. %v\n", x, i)
		case 3:
			fmt.Printf("%d: Three | Iteration no. %v\n", x, i)
		case 4:
			fmt.Printf("%d: Four | Iteration no. %v\n", x, i)
		default:
			fmt.Printf("%d: Unknown Number | Iteration no. %v\n", x, i)
		}
	}
}
