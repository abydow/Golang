package main

import (
	"fmt"
)

func main() {
	for j := range 100 {
		if j%2 != 0 {
			fmt.Printf(" %v |", j)
			continue
		}
	}
	fmt.Println("")
}
