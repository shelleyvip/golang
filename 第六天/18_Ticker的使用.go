package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second*1)
	i := 0
	for{
		<-ticker.C
		i ++
		fmt.Println("i=",i)
		if i == 5{
			ticker.Stop()//到5的时候 stop
			break  //还有一个操作就是跳出这个循环
		}
	}
}

