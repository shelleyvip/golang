package main

import (
	"fmt"
	"strconv"
	"time"
)

//在主线程（可以理解成进程）中，开启一个goroutine,该协程每隔1秒输出"hello world"
//在主线程中也每隔一秒输出"hello golang",输出10次后退出
//要求主线程和goroutine一起执行



func test()  {
	for i := 1; i <=10;i++{
		fmt.Println("test() hello world"+ strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main()  {
	go test()
     for i :=1;i<=10;i++{
     	fmt.Println("main()hello golang"+strconv.Itoa(i))
     	time.Sleep(time.Second)
	 }
}
