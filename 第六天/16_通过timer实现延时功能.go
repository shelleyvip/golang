package main

import (
	"fmt"
	"time"
)

////实现timer延时功能共有以下三种方法，均可达到目的
func main(){
	//常用
	<- time.After(time.Second*2)//定时2s，阻塞2s,2后产生一个事件，往timer里面写内容
	fmt.Println("时间到")
}
func main02(){
	time.Sleep(time.Second*2) //在或者是让它睡上2s 在打印
	fmt.Println("时间到")
}
func main01(){
	//延时2s后打印一句话
	timer := time.NewTimer(time.Second*2)
	<-timer.C
	fmt.Println("时间到")
}