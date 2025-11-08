package main

import "fmt"

const (
	_ = iota
	a
	b
	c
	d
	e
	f
)

func main() {

	fmt.Printf("%b in binary and %d in decimal\n", 1<<a, 1<<a)
	fmt.Printf("%b in binary and %d in decimal\n", 1<<b, 1<<b)
	fmt.Printf("%b in binary and %d in decimal\n", 1<<c, 1<<c)
	fmt.Printf("%b in binary and %d in decimal\n", 1<<d, 1<<d)
	fmt.Printf("%b in binary and %d in decimal\n", 1<<e, 1<<e)
	fmt.Printf("%b in binary and %d in decimal\n", 1<<f, 1<<f)

}
