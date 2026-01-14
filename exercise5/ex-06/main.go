package main

import "fmt"

func main() {
	fmt.Println("Exercise 6")
	xs := make([]string, 0, 51)
	fmt.Println(len(xs))
	fmt.Println(cap(xs))
	fmt.Println("----------------------------")

	xs = append(xs, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println(len(xs))
	fmt.Println(cap(xs))
	for i, v := range xs {
		fmt.Printf("%v\t%T\t%v\n", v, v, i)
	}

}
