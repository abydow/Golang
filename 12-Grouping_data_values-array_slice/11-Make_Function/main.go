package main

import (
	"fmt"
)

func main() {
	si := []string{"A", "B", "C"}
	fmt.Println(si)

	sf := make([]int, 0, 10)
	fmt.Println(sf)
	fmt.Println(len(sf))
	fmt.Println(cap(sf))
	sf = append(sf, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("%v \n",sf)
	fmt.Println("--------------------------")
	fmt.Println(len(sf))
	fmt.Println(cap(sf))
	fmt.Println("--------------------------")
	sf = append(sf, 11, 12, 13, 14, 15)
	fmt.Println(len(sf))
	fmt.Println(cap(sf))
	fmt.Printf("%v \n",sf)

}
