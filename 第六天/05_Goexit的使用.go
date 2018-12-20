package main

import (                   ////Goexit():就是结束当前协程的
	"fmt"
	"runtime"///runtime.Goexit()  中止所在的协程// 所在的协程也就是说所有的协程
)

func test(){

	defer fmt.Println("bbbbbbbbb") //3）但是这个defer 函数是退出之前执行的，所以只有它可以执行
	//return   // 1）:中止此函数下面的内容，不执行了 返回的意思
	runtime.Goexit() //2）:中止所在的协程// 所在的协程也就是说所有的协程
	fmt.Println("ccccccccc")
}


func main(){
	//创建一个新的协程
	go func(){
		fmt.Println("aaaaaaaaaa")
		test()
		fmt.Println("ddddddd")

	}()//别忘记后面的（）

	for{
          //特地一个死循环：目的是不让主协程结束
	}
}
