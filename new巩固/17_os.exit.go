package main

import (
	"fmt"
	"os"
)

func print()  {
	fmt.Println("hello")
	
}
func main() {
	_,err := os.Open("a.txt")
	if err != nil{
		fmt.Println(err)
		os.Exit(2)
		return
	}
	print()
}
