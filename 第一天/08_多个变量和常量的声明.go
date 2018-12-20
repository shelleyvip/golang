package main

import "fmt"

func main() {
	//不同变量类型的声名
	var a,b = 10,3.14
	fmt.Printf("a=%T,b=%T\n",a,b)

//不同常量类型的声明
	const (
		c  int = 10
		e float64 = 3.14
	)
	fmt.Println("c=", c)
	fmt.Println("e=", e)

}