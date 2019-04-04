package main

import (
	"fmt"
	"runtime"
)

func test()  {
	defer fmt.Println("cccccccccc")
	//return //终止了 也就是说返回了此函数不再继续了
	runtime.Goexit() //这个函数意思是终止所在的协程 就是说整个协程就结束了

	fmt.Println("dddddddddd")

}

func main() {
	//创建一个新的协程
	go func() {
		fmt.Println("aaaaaaaaaaaa")
		//调用别的函数
		test()
		fmt.Println("bbbbbbbbbbb")

	}()

	for{

	}
}
