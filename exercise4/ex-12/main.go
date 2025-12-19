package main

import (
	"fmt"
)

func main() {
	for i := range 5 {
		for j := range 5 {
			fmt.Printf("i: %d , j: %d\n", i, j)
		}
	}
}
