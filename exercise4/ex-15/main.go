package main

import (
	"fmt"
)

func main() {
	x := map[string]int{
		"Legend":   32,
		"Zane": 27,
		"Adi":       64,
	}
	// m1 := x["Adi"]
	// fmt.Println(m1)

	if v, ok := x["Adi"]; ok {
		fmt.Println("The value is:", v)
	}; if v, ok := x["John"]; ok {
		fmt.Println("The value is:", v)
	}; if v, ok := x["Legend"]; ok {
		fmt.Println("The value is:", v)
	}else {
		fmt.Println("The key does not exist.")
	}

}
