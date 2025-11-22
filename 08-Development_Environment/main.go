package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)   //shows Operating System
	fmt.Println(runtime.GOARCH) // shows system architechture

}

// go run ./...
// go build (turns it to executable)
// to build for a specific system we have to use GOOS = darwin(mac)/windows/linux go build
