package main

import "fmt"
var a int
//什么全局变量？ 答：在函数以外的变量
//全局变量的 作用域 是所有函数

func test()  {

	fmt.Println("a=",a)
}

func main(){
	a = 88
	test()
	fmt.Println("a=",a)
}


