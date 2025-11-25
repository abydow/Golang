package main

import "fmt"

func dep1() {
	fmt.Println("I am inside dep1.go *Direct dependency* ")
	dep_2()
}
