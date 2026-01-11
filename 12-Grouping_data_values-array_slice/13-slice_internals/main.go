package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("There is a array lying underneath slices!")
	a:= []int{1, 2, 3, 4, 5}
	b:= make([]int, 6) // create diff. underlying array
	copy(b, a)
	fmt.Println("---------------------------")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("---------------------------")

	a[0] = 7
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("---------------------------")



//-------------------------------------------

x:= []float64{3,1,4,2}
n:= []float64{3,1,4,2}
fmt.Println(medianOne(x))
fmt.Println(x)
fmt.Println(medianTwo(n))
fmt.Println(n)

}

func medianOne(x []float64) float64 {
	sort.Float64s(x)
	i := len(x) / 2
	if len(x)%2 == 1 {
		return (x[i/2])
	}
	return (x[i-1] + x[i]) / 2

}
func medianTwo(x []float64) float64 {
  n:= make([]float64, len(x))
  copy(n, x)

  sort.Float64s(n)
  i := len(n) / 2
  if len(n)%2 == 1 {
    return (n[i/2])
  }
  return (n[i-1] + n[i]) / 2
}
