package main

import "fmt"

func swap(x, y string) (string, string /*we used this to tell the console that there will be two values that will return and both will be string*/) {
	return y, x
}

func main() {
	a, b := swap("dan", "dadan")
	fmt.Println(a, b)
}
