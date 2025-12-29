package main

import (
	"fmt"
)

func main(){
	fmt.Println("Demonstrating access by position using for range:")
	xf := [] string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}
	for _, v := range xf {
		fmt.Printf("%s\n",v)
	}

// To print by index position>>
	fmt.Println("\n --------------------")
	fmt.Println(xf[0])
	fmt.Println(xf[1])
	fmt.Println(xf[2])
	fmt.Println("\n --------------------")
	fmt.Println(len(xf))
	fmt.Println("\n --------------------")
	xl := []int{42, 43, 44, 45, 46}
	for i := 0; i < 3; i++ {
		fmt.Println(xl[i])
	}
}
