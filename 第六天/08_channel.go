package main

import "fmt"

////定义一个全局变量，创建一个channel
var ching = make(chan int)
//Public ： 公共
func Public(str string)  {
	for _,v := range str{
		fmt.Printf("v=%v\n",v)
	}
	
}
////这里目的是让Senduot1执行完之后，在执行Receive
func Senduot1()  {
//	fmt.Println("Senduot发送的意思")
	Public("Senduot发送的意思")
	 ching<- 999  //往通道发送99的数据
}
//Receive 接收的意思
func Receive()  {
    <-ching     //从管道里取数据，接收，如果通道里没有数据，他就会阻塞
	Public("接收")


}


func main()  {
	go Senduot1()
	go Receive()
	for {

	}


}
