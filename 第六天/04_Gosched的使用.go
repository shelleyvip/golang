package main

import (
	"fmt"
	"runtime"  //runtime.Gosched 是主协程让出时间片让子协程执行完毕之后再回来执行此协程
)

func main() {
	go func() {
		for i :=0;i<10;i++{
			fmt.Println("go")
		}


	}()
    //在这里我的主协程的任务执行完了就结束了
	for i := 0; i < 2;i++{
		runtime.Gosched() //<-所以在这里让出时间片,让其它子协程执行完毕之后，在回来执行此协程
		fmt.Println("hello")  //runtime.Gosched()让出时间片
	}
}