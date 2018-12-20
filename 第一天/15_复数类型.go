package main

import "fmt"

func main(){
	var a complex128 //声明复数类型
	a = 3.3+3.14i //3.3实部，3.14i是虚部
	fmt.Println("a=",a)

/////自动推导类型
	zhou := 3.3 + 3.14i
	fmt.Printf("zhou=%T\n",zhou)

	//通过内建函数取实部和虚步
	fmt.Println("real(zhou)=",real(zhou),"imag(zhou)=",imag(zhou))
}
