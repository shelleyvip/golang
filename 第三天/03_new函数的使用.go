package main

import "fmt"

func main()  {
	a := new(int)
	*a = 666
	fmt.Println("*a=",*a)

}
