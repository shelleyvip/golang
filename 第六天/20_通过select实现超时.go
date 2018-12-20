package main

import (
	"fmt"
	"time"
)
func main()  {
	ch := make(chan int)
	quit :=make(chan bool)
	go func() {
		for {
			select {
			case num:= <-ch:
				fmt.Println("num=",num)
			case <-time.After(time.Second*3):
				fmt.Println("超时了......")
				quit<-true
			}
		}
	}()
	for i:=0;i<6;i++{
		ch<- i
		time.Sleep(time.Second)
	}
	<-quit
	fmt.Println("程序结束🔚")
}