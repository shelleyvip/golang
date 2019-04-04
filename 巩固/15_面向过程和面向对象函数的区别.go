package main

import "fmt"

type loog int

func  (tmp loog)Add(other loog)loog{
	return tmp + other

}
func main() {
	var  a loog = 2
	result := a.Add(3)
	fmt.Println(result)

}
