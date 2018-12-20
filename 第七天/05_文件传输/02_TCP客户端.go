package main

import (
	"fmt"
	"net"
)

//使用方法：/先提供服务，在用 run 和build 都可以（适用于所有的程序）
//先运行tcp服务端，在运行客户端
func main(){
	//主动连接服务器
	conn,err := net.Dial("tcp","127.0.0.1:8000")
	if err != nil{
		fmt.Println("err=",err)
		return
	}
	defer conn.Close()
     //发送数据
	conn.Write([]byte("are you ok?"))
}