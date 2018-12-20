package main

import "fmt"

func main(){
	//变量var:是在程序运行期间是可以改变的量
	//声明好的常量和类型
	const a int = 10     //常量const:在程序运行期间是不可以改变的量
	fmt.Println("a=",a)


	//自动推导类型
	const b = 88
	fmt.Printf("b= type is %T\n",b)
	//打印还可以直接使用%T，结果是一样的
	fmt.Printf("b=%T\n",b)
}