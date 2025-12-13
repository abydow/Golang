package main

import "fmt"

func main(){
	fmt.Println("")
	fmt.Println("Prime numbers between 0 and 100 are:")
	fmt.Println("-----------------------------------")
	fmt.Printf("1 is a prime number\n")
	fmt.Printf("2 is a prime number\n")
	for i:=0; i<=1000; i++{
		for j:=2; j<i; j++{
			if i%j == 0{
				break
			}
			if j == i-1{
				fmt.Printf("%v is a prime number\n", i)
			}
		}
	}

}
