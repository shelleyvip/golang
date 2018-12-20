package main

import (
	"fmt"
	"time"
)

func Newtask(){
	for  {
		fmt.Println("子携程")
		time.Sleep(time.Second*3)//延时1S 一秒

	}
}

func main()  {
	go Newtask() //新建一个协程，仙剑一个新任务
	for  {
		fmt.Println("主协程")
		time.Sleep(time.Second*3)
	}


}
