package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello, go!")
	fmt.Print(runtime.GOARCH)
}
