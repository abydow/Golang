package main

import "fmt"

func main() {

	x := 42

	if x > 20 {
		fmt.Println("use of greater than function")
	}

	if x == 42 {
		fmt.Println("use of ==")
	} else if x >= 42 {
		fmt.Println("use of >=")
	} else {
		fmt.Println(" learned how to write a if-else_if-else block") // will not execute cause the if function woks
	}
	/*
		according to control flow :--

		if < condition > {} if the condition meets rest of the code will not work
		else if < condition > {} if the if function didnt ran it will run.
		else if < condition > {} only runs if the previous func. didnt work.
		else {} only run if no conditions were met
	*/
	if x < 40 {
		fmt.Println("It will not run, x is not less than 40 ")
	} else if x <= 42 {
		fmt.Println("It will run . x is equal to 42")
	} else if x != 40 {
		fmt.Println("Learned not equal to sign")
	}
}

// <= and >= are like or function with two condition
// "a<=b" = a<b or a==b
// "a>=b" = a>b or a==b
