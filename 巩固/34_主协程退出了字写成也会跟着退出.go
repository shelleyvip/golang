package main

import (
	"fmt"
	"time"
)

func main() {


	go func() {
		i :=0
		for{
			i++
			fmt.Println("子协程")
			time.Sleep(time.Second)
		}
	}()


	i := 0
	for{
		i++
		fmt.Println("主协程")
		time.Sleep(time.Second)  //猪协程退出了 子协程也会跟着退出
		if i == 2{
			break
		}
	}
}
