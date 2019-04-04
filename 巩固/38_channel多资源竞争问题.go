package main

import (
	"fmt"
	"time"
)
//创建一个channel全局变量 根据需要要选择类型
var ch = make(chan int)

//这个相当于一个打印机是属于公共资源
func Println(str string){
	for _,data := range str{
		fmt.Printf("%c",data)
		time.Sleep(time.Second)
	}
fmt.Printf("\n")
}
//Person1执行完后，才能执行到Person2
func Person1(){
Println("hello")
   ch <- 666 //给管道写数据，发送
}

func Person2(){
	<- ch //从管道里取数据 接收 如果通道没有数据就会阻塞
	Println("world")

}

func main() {
	//新建两个协程相当于两个人同时使用打印机
	go Person1()
	go Person2()

	//特意写个死循环 目的是不让主协程结束
	for{

	}

}


