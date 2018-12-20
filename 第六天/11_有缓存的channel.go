package main

import (
	"fmt"
	"time"
)
var ch = make(chan int)
func main()  {
	go func() {
		for i:=0;i<100;i++  {
			ch<-i
			fmt.Printf("len=%v cap=%v i=%v\n",len(ch),cap(ch),i)


		}
	}()
	time.Sleep(time.Second*2)
	for i:=0;i<100;i++{
		  <-ch
		fmt.Printf("i=%v",i)
	}

}