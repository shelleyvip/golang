package main

import "fmt"

func main() {
   //创建一个无缓存的管道channel
	ch := make(chan int)

	//新建一个goroutine
	go func() {
		for i:=0;i<5 ;i++  {
			ch <-i //往通道里写数据
		}
		//不需要在写数据时，关闭channel
		close(ch)
		//ch <-666 //关闭通道后就没法在发送数据
		//但可以接收数据

	}()

	for{
		//如果ok为true 说明管道没有关闭
		if num,ok := <-ch;ok == true{
			fmt.Println("num=",num)
		}else { //管道关闭
			break
		}
	}
}
