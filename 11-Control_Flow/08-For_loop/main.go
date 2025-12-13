package main

import "fmt"

func main(){

	x := 0

	for i:=0; i<=5; i++ {    // for condition; increment/decrement statement;
		fmt.Printf("Iteration number: %v\n", i)
	}

	for x<10 {
		fmt.Printf("Value of x: %v\n", x)
		x ++
	}
	// break statement
	z := 0
	for {
		fmt.Printf("loop until 16 %v\n", z)
		if z > 16 {
			break // completely exists the loop
		}
		z=z+2
	}
	// continue statement
	for j:=range 10 { // continues even if the if statement isn't met
		if j%2 == 0 {
			fmt.Printf("even number: %v\n", j)
			continue
		}
	}
}
