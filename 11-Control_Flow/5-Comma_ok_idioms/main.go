package main

import (
	"fmt"
)
func main(){
	x := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}

	if num, ok := x["one"]; ok {      // so when we are using two variables num, ok
		fmt.Printf("%d %v \n", num, ok) // the program will return the value and boolean
	}			                            // cause in go if the value exists or not it returns 0
	                                  // so when we return two values we can check if the value exists or not
	                                  // by checking the boolean value
	                                  // this is called comma ok idiom
	                                  // it is mostly used in maps but there are other use cases too
	if num, ok := x["two"]; ok {      // we will explore that later
		fmt.Printf("%d %v \n", num, ok)
	}
	if num, ok:= x["three"]; ok {
		fmt.Printf("%v %v\n", num, ok)
	} else {
		fmt.Println("not found")
	}
}
