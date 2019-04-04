package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个channel通道
	ch := make(chan string)
	defer fmt.Println("主协程也结束了")


	go func() {
		defer fmt.Println("子协程调用完毕")
		for i:=0;i<2;i++{
			fmt.Println("i=",i)
			time.Sleep(time.Second)
		}
         ch <- "我是子协程 我工作完毕！"
	}()


	str := <- ch   //没有数据前阻塞
	fmt.Println("str=",str)


}
