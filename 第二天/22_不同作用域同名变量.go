package main

import "fmt"

var a int
func main(){

	fmt.Printf("a= %T\n",a)
	{
		var a byte
		fmt.Printf("a=%T\n",a)
	}

	{
		var a string
		fmt.Printf("a=%T\n",a)
	}
	test()
}
func test()  {
	fmt.Printf("%T\n",a)

}

