package main

import "fmt"

func Pritin() {
	defer func() {
		fmt.Println("defer")
	}()
	fmt.Println("hello")
}

func main() {
	Pritin()
}
