package main

import "fmt"

const Pi float32 = 3.14

func main() {
	const world = "Duniya"
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
