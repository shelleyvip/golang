package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(1) //指定以多少核来计算

	fmt.Println("n=",n)

	for{
		go fmt.Print(6)
		fmt.Print(8)
	}

}
