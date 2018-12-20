package main

import "fmt"

func main()  {
	ch := make(chan int)
	//新建一个goroutine
	go func() {
		for i:=0;i<5;i++{
			ch<-i
		}
		close(ch)//当不需要在往goroutine写东西的时候，关闭
	}()
	for num :=range ch{      //这是比较简便的一种写法
		fmt.Println("num=",num)
	}
	for {  //这是第二种写法
		if nem,ok := <-ch;ok ==true{
			fmt.Printf("nem=%v",nem)
		}else {
			break //
		}
	}
}