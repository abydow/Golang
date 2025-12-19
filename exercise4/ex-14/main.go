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

	m1 := x["Zane"]
	fmt.Println(m1)

	}
