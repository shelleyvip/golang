package main

import (
	"fmt"
	"time"
)
//验证Timer.NewTimer()时间到了之后响应一次
//func mai01(){
//	timer :=time.NewTimer(time.Second*1)
//	for{
//		<- timer.C
	//	fmt.Println("时间到")
	//}
//}

func main()  {

	//在这里创建一个定时器，设置时间为2S，2s后往time通道写内容（当前时间）
	timer := time.NewTimer(time.Second*2)
	fmt.Println("当前时间:",time.Now())
	//2s后，往timer.c写数据，有数据后，就可以读取
	t := <- timer.C
	fmt.Println("t=",t)

}
