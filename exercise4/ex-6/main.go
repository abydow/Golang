package main

import (
	"fmt"
	"math/rand"
)
func main() {
	for i:=0;i<5;i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Printf("x: %d , y: %d\n", x, y)
		comswitch(x, y)
	}
}
// if x and y both less than 4
// if x and y both greater than 6
//if x >=4 and <=6
//y != 5
//none of the above


func comswitch(x,y int)  {
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
