package main

import (
	"fmt"
)

func main()  {
	//创建一个无缓存的channel
	ch := make(chan int)   //len缓冲区剩余数据个数，cap 缓冲区大小，上面是无缓冲区所以是0
	fmt.Printf("len(ch)=%v cap(ch)=%v\n",len(ch),cap(ch))
	go func() {
		for i :=0;i<5;i++ {
			fmt.Println("i=子协程执行完毕",i)
			ch<-i


		}
	//	fmt.Println("i=子协程执行完毕")

	}()
	//time.Sleep(time.Second*2)
	for  i:=0;i<5;i++ {
		  <-ch  //开始读取管道中的内容，若没有内容，堵塞
		fmt.Println("主协程执行完毕")

	}

}
