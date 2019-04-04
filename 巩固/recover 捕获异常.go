package main

import (
	"fmt"
	"time"
)

func sayhello()  {
	for i:= 0;i<10;i++{
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}


func test1()  {

	defer func() {
		if err:= recover();err !=nil{
			fmt.Println("test1()发生错误",err)  //捕获错误
		}
	}()
   //定义一个map
   var myMap map[int]string
   myMap[0] = "golang"
}

func main() {
	go sayhello()
	go   test1()

	for i:=0;i<10;i++{
		fmt.Println("main()ok=",i)
		time.Sleep(time.Second)
	}


}
