package main

import "fmt"

func main() {
	s := make([]int,0,1)
	s = append(s,1)
	fmt.Println(s)

	s = append(s,2)
	fmt.Println(s)
}
