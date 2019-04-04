package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个有缓存放入channel
	ch := make(chan int, 3)
	fmt.Printf("len(ch)=%d,cap(ch)=%d\n", len(ch), cap(ch))


	//新建一个子协程
	go func() {
		for i :=0;i<10;i++{
			ch <- i
			fmt.Printf("子协程:[%d] len(ch)=%d,cap(ch)=%d\n",i,len(ch),cap(ch))
		}
	}()
	//延时
	time.Sleep(2*time.Second)
	for i:=0;i<10;i++{
		num := <-ch
		fmt.Println("num=",num)

	}
}
