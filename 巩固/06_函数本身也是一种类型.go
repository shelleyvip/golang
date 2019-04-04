package main

import "fmt"

type FuncType func(int,int)int

func add(a,b int)int  {
	return a + b
}
func minus(a,b int)int  {
	return a - b
}
     //首先前提是调用的类型参数和返回值都是一样的 情况下可以调用
func main() {
	var fTest FuncType  //生命一个函数类型的变量
	fTest = add  //把add函数的值给ftest
	result := fTest(10,20) //生命一个接收者 给ftest赋值

	fmt.Println("result=",result)



}
