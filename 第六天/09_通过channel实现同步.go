package main

import "fmt"

func main()  {
	var ch  = make(chan string)
	defer fmt.Println("这是主协程")
	go func() {
		for i:=1;i<6;i++ {
			fmt.Println("这是子协程 i=",i)
        // ch<-"周雪莉"
		}

	}()

	//主协程
	 str := <-ch //读取ch管道的数据，并且把变量传给str
	 fmt.Println("str",str)
}
