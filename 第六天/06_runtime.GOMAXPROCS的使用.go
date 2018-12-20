package main

import (
	"fmt"
	"runtime"
)

func main()  {
	//n := runtime.GOMAXPROCS(1)//指定CPU的核心数 1 就是代表1核处理
	n := runtime.GOMAXPROCS(4)//就代表4核
	fmt.Println("周雪莉",n)
	for {
		fmt.Print(0)//用print 是直接打印数字，也就是说你输入什么打印什么
	go 	fmt.Print(1)
	}
}
