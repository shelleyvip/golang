package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {   //新开一个协程
		for{      //不停的检测数据的流动
			select {
			case num := <-ch://把管道赋值给num 如果这个num有数据 就开始打印
				fmt.Println("num=",num)
			case <-time.After(3*time.Second): //如果长时间没有数据 延时三秒钟
			fmt.Println("超时")
				
			quit <- true
				
			}
		}
	}()
	for i := 0;i<5;i++{
		ch <- i //把i传给ch
		time.Sleep(time.Second)//延时一秒打印
	}
	<- quit
	fmt.Println("程序结束")
}
