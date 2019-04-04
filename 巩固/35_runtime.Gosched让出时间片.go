package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i :=0;i<5;i++{
			fmt.Println("go")
		}
	}()

	for i :=0;i<2;i++{
		runtime.Gosched() //让出时间片 让其他程序执行完毕之后 在回来执行此程序
		fmt.Println("hello")
	}

}
