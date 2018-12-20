package main

import (
	"fmt"
	"time"
)

func main()  {
////主协程退出了其它子协程也跟着退出
	go func(){
	i := 0
		for {
			i++
			fmt.Println("子协程:i=", i)
			time.Sleep(time.Second)
		}

	}()//别忘了()

	i := 0
	for{
		i++
		fmt.Println("主协程:i=",i)
		time.Sleep(time.Second)
		if i == 2{
			break
		}
	}

}
