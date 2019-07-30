package main

import "fmt"

func main() {
	intChan := make(chan int,3)
	intChan <- 10
	intChan <- 20
	close(intChan)//关闭管道

	//当关闭管道后不能在继续写入
	//intChan <- 30

	fmt.Println("ok ook")

	//当管道关闭后 读取数据是可以的
	n := <-intChan
	fmt.Println("n=",n)
}
