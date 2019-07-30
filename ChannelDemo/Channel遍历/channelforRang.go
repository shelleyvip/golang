package main

import "fmt"

func main() {
	//遍历管道
	intChan := make(chan int,100)
	for i:=0;i<100;i++{
		intChan <- i*2  //放入100个数据到管道
	}
	//遍历管道不能用普通的for循环
	//for i :=0;i<len(intChan);i++{
	//	error 错误
	//}

	//1:要用for range
	//2:在遍历时channel管道没有关闭 则会出此案deadlock的错误

	for v := range intChan{
		fmt.Println("v=",v)  //管道是没有下标的 是队列
	}

}
