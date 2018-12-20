package main

import "fmt"

func main(){
	var a bool    //声明变量 没有初始化 初始值为false
	a = true
	fmt.Println("a=",a)

	var b false   ////自动推导
	fmt.Println("b=",b)

	c := false
	fmt.Println("c=",c) ///// 简单声明



}