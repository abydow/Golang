package main

import (
	"fmt"
	"math/rand"
)

func init () {
	for i := 0; i <= 15; i++ {
		x:= rand.Intn(250)
		fmt.Printf("x:-> %d  \n", x)
	}
}

func main() {
	fmt.Println("See init function in action!")

}
