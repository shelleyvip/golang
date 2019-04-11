package main

import "fmt"

func sum(args ...int)int  {
	n := 0
	for i :=0;i<len(args);i++{
		n += args[i]
	}
	return n
	
}

func main() {
	fmt.Println(sum(1,2,3,4,5,6,7,8,9))
	s := []int{1,2,3}
	fmt.Println(sum(s...))
}
