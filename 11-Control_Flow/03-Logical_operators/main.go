package main

import "fmt"

func main() {
	a := 13
	b := 12

	if a > b && b < a {
		fmt.Println("This block will only run if both of the conditions were met")
		// here a is greater than b
		// which also implies b is greater than a
		// so the line will print

	}

	if a > b || b > a {
		fmt.Println("runs if any one condition is true")
		// in or function only one condition need to be fullfilled
		// if the 1st condition is met the code will run
		// if its not, the second condition will be checked
		// if it meets the code will run or else it wouldnt
		// here a>b is correct but b>a is not. but due to satisfaction of one condition the code will run
	}

	if a != b {
		fmt.Println("only runs if the condition is not satisfied")
		// it is much simple to understand. if the condition is not true the code will run . for example :=
		// a is not equal b
		// so we used the "!" this not operator to put a condition
		// if a is not equal to b print "....."
	}
}
