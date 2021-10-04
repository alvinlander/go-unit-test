package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.Compiler)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOROOT())
}
