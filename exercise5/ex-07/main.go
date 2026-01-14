package main

import "fmt"

func main() {
	fmt.Println("Exercise 7")

	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "I'm 008."}

	xp := [][]string{a, b}

	for i, v := range xp {
		fmt.Println(v,i)
	}
}
