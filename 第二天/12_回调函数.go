package main

import "fmt"

type Shelley01 func(int,int)int

//实现加法
//多态，多种形态调用同一个接口，不同的表现，可以实现不同的表现，加减乘除
func Zhou(a,b int)int{
	return (a + b)
}
//回调函数，函数有一个参数是函数类型 这个函数就是回调函数
//计算机可以实现四则运算
//先有想法后面在实现功能
func hello(a,b int,nihao Shelley01)(milu int){
	fmt.Println("hello=",hello)
	milu = nihao(a,b)
	return
}

func main(){
	abc := hello(90,100,Zhou)
	fmt.Println("abc=",abc)

}
