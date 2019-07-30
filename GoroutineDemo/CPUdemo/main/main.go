package main

import (
	"fmt"
	"runtime"
)

func main() {
	//获取当前系统CPU的数量
	numcpu := runtime.NumCPU()
	fmt.Println("numcpu=",numcpu)

	//可以自己设置使用多个CPU
	runtime.GOMAXPROCS(numcpu - 1)
}




