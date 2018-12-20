package main

import (
	"fmt"
	"time"
)
//重置定时器
//Reset :重置的意思
func main(){
	timer := time.NewTimer(time.Second*5)
	timer.Reset(time.Second*1)
	<-timer.C
	fmt.Println("时间到")

}
//停止定时器
func main01() {
	timer := time.NewTimer(time.Second*3)
	go func() {
		<- timer.C
		fmt.Println("子协程可以打印了，因为定时器的时间已经到了")
	}()
   timer.Stop() //在这里停止定时器     //那么就goroutine 的内容就打印不出来了
	for {

	}
}

