package main

import "fmt"
//无参无返回值的调用
func Myxueli(){ //// Myxueli 自定义声明一个函数名
	a := 666          //给它赋个值
	fmt.Println("a=",a) //把这个值打印出来
}
func main(){          ///1：首先程序必须从main 包开始执行
	Myxueli() ////  然后 调用Myxueli（） 里面的值
}
