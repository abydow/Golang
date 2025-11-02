package main

import "fmt"

func main() {

	// when we use casting() to convert one data type to another It is called Type casting
	// conversion allows the compiler to convert one data type to another (in dynamic languages)

	// for example

	y := 24
	z := 24.0

	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	var m float32 = 43.742
	fmt.Printf("%v of type %T \n", m, m)

	/*
		#In go you cant assign a float32 value to a variable that is holding float64

		like;

		z=m  // ❌
		fmt.Printf("%v of tye %T \n", z, z)


	*/

	// If we need to assign convert it first

	z = float64(m)
	fmt.Printf("%v of type %T \n", z, z)

}
