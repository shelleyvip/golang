package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个无缓存的channel
	ch := make(chan int,0)
	fmt.Printf("len(ch)=%d,cap(ch)=%d",len(ch),cap(ch))
	//新建一个协程
	go func() {
		for i :=0;i<3;i++{
			fmt.Printf("子协程 i =%d",i)
			ch <- i  //往管道里内内容
			//fmt.Printf("len(ch)=%d,cap=(ch)=%d\n" ,len(ch),cap(ch))
		}
	}()

	time.Sleep(2*time.Second)
	for i := 0;i<3;i++{
		num := <- ch //读管道的内容，没有内容阻塞
		fmt.Println("num=",num)
	}


}
