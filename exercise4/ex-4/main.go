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
	fmt.Println(`
Niladic (0 parameters) - func doSomething()
Monadic (1 parameter) - func square(x int)
Dyadic (2 parameters) - func add(a, b int)
Triadic (3 parameters) - func volume(l, w, h int)
Polyadic (many parameters) - func sum(nums ...int)
	`)

}
