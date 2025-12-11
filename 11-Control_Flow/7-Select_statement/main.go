package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received", msg1, "from ch1")
	case msg2 := <-ch2:
		fmt.Println("Received", msg2, "from ch2")
	}

}
