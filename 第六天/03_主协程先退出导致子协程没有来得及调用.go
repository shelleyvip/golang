package main

import (
	"fmt"
	"time"
)

func main()  {
   //主协程退出了 所以子协程就没办法在继续调用了
	go func(){ //创建一个子协程
		i := 0
		for {
			i++
			fmt.Println("子协程:i=", i)
			time.Sleep(time.Second)
		}
  //这是在同一个main下面的内容，所以主协程和子协程哪个限制性，这个要看系统的时间戳 美欧顺序而言

	}()//别忘了（）

	i := 0
	for{  //这里是主协程
		i++
		fmt.Println("主协程:i=",i)
		time.Sleep(time.Second)
		if i == 2{  //在这里主协程的命令就执行到<-2 就跳出来了，所以没有办法在去执行子
			break/////协程的命令了，所以叫主协程退出子协程没有来得及调用
		}
	}

}

