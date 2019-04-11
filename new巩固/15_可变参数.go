package main

import "fmt"

func sum(args ...int)int  {
	n := 0
	for i := 0;i<len(args);i++{
		n += args[i]
	}
	return n
}

func main() {
	fmt.Println(sum(3))
	s := []int{3}
	fmt.Println(sum(s...))
}
