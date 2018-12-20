package main

import "fmt"

func zhou(a int){
	if a == 1{
		fmt.Println("a=",a)
		return
	}
	zhou (a - 1)
	fmt.Println("a=",a)
}
func main (){
	zhou(3)
	fmt.Println("hello")
}