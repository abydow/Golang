package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i:=0;i<=10;i++ {

		x := rand.Intn(250)
		fmt.Printf("x: %d -> ", x)

		if x < 100 {
			fmt.Println("Between 0 and 100")
		}
		if 101 < x && x < 200 {
			fmt.Println("Between 100 and 200")
		}
		if 201 < x && x < 250 {
			fmt.Println("Between 200 and 250")
		}
	}

	for i:=0;i<=15;i++ {
		fmt.Printf(" %d\t", rand.Intn(2)) // binary digits
	}
	fmt.Println("")
}
