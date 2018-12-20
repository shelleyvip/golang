package main

import (
	"fmt"
	"net"
)

func main(){
	//主动连接服务器
	conn,err := net.Dial("tcp",":8000")
	if err != nil{
		fmt.Println("net.Dial err",err)
		return
	}
	defer conn.Close()
	requestBuf := "GET/go HTTP/1.1\r\nAccept:image/gif,image/jpeg,image/pjpeg,application/x-ms-application,application/xaml+xml,application/x-ms-xbap,*/*\r\nAccept-Language:zh-hans-CN,zh-hans;q=0.8,en-US;q=0.5,en;q=0.3\r\nUser-Agent:Mozilla/4.0(compatible;MSIE 7.0;Windows NT 10.0;WOW64;Trident/7.0;.NET4.0C;.NET4.0E;.NET CLR 2.0.50727;.NET CLR3.0.30729;.NET CLR 3.5.30729)\r\nAccept-Encoding:gzip,deflate\r\nHost:127.0.0.1:8000\r\nConnection:Keep-Alive\r\n\r\n"

	//先发送请求包，服务器才会回响应包
	conn.Write([]byte(requestBuf))

	//接收服务器回复的响应包

	buf := make([]byte,1024*4)

	n,err := conn.Read(buf)
	if err != nil{
		fmt.Println("conn.read err",err)
		return
	}
	//打印响应包文
	fmt.Printf("#%v#",string(buf[:n]))


}
