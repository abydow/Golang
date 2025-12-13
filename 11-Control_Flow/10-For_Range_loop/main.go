package main

import "fmt"

func main(){
	xi := []int{10,20,30,40,50}

	for i,v := range xi{
		fmt.Printf("index: %v, value: %v\n", i, v)
	}

	z := map[string]int{
		"apple":  10,
		"banana": 20,
		"cherry": 30,
	}

	for k,v := range z{
		fmt.Printf("key: %v, value: %v\n", k, v) // randomized order
	}
}
