package main

import "fmt"

func add(x, y int /*we can also write (x int, y int)*/) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
	sayHello()
}

func sayHello() {
	fmt.Println("Hello")
}
