package main

import "fmt"

func main(){
	var a int //赋值前必须先声明变量
	a = 10    //赋值可以使用N次
	a = 20
	a = 30
	fmt.Println("a=",a)
   // := 是自动推导类型
   //自动推倒 同一个变量名只能使用一次 用于初始化那次
   b := 88    //在这里已经有了一个变量b，不能在使用变量b
   fmt.Printf("b type is %T\n",b)
}
