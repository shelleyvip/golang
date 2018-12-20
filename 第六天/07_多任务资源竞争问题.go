package main

import (
	"fmt"
	"time"
)
//定义一个打印机，参数为字符串，按每个字符打印
//打印机属于公共资源
func Printer( str string){
	for _,data := range str {
		fmt.Printf("%c",data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}
func Prinson1(){
	Printer("hello")
}
func Prinson2(){
	Printer("world")
}



func main(){
	//在这里新建两个协程，代表两个人，两个人同时使用打印机
	go Prinson1()
	go Prinson2()
	for{
///在这里来个死循环，特地不让主协程结束
	}
}