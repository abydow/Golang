package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for j := range 100 {
		fmt.Printf("%v ", j)
	}
	fmt.Println("")
	for i := 0; i < 100; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)
		fmt.Printf("x: %d , y: %d\n", x, y)

		switch {
			case x < 4 && y < 4:
				fmt.Println("x and y both less than 4")
			case x > 6 && y > 6:
				fmt.Println("x and y both greater than 6")
			case x >= 4 && x <= 6:
				fmt.Println("x is between 4 and 6")
			case y != 5:
				fmt.Println("y is not equal to 5")
			default:
				fmt.Println("none of the above")
			}
	}
}
