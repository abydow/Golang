package main

import "fmt"

func main(){
	for i:=0; i<=5; i++ {
		for j:=0; j<=5; j++ {
			fmt.Printf("outer loop: %v \t inner loop: %v\n", i, j)
		}
	}
}
/*
Output:
outer loop: 0    inner loop: 0
outer loop: 0    inner loop: 1
outer loop: 0    inner loop: 2
outer loop: 0    inner loop: 3
outer loop: 0    inner loop: 4
outer loop: 0    inner loop: 5
outer loop: 1    inner loop: 0
outer loop: 1    inner loop: 1
outer loop: 1    inner loop: 2
outer loop: 1    inner loop: 3
outer loop: 1    inner loop: 4
outer loop: 1    inner loop: 5
outer loop: 2    inner loop: 0
outer loop: 2    inner loop: 1
outer loop: 2    inner loop: 2
outer loop: 2    inner loop: 3
outer loop: 2    inner loop: 4
outer loop: 2    inner loop: 5
outer loop: 3    inner loop: 0
outer loop: 3    inner loop: 1
outer loop: 3    inner loop: 2
outer loop: 3    inner loop: 3
outer loop: 3    inner loop: 4
outer loop: 3    inner loop: 5
outer loop: 4    inner loop: 0
outer loop: 4    inner loop: 1
outer loop: 4    inner loop: 2
outer loop: 4    inner loop: 3
outer loop: 4    inner loop: 4
outer loop: 4    inner loop: 5
outer loop: 5    inner loop: 0
outer loop: 5    inner loop: 1
outer loop: 5    inner loop: 2
outer loop: 5    inner loop: 3
outer loop: 5    inner loop: 4
outer loop: 5    inner loop: 5
*/
