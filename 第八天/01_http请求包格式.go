package main

import (
	"fmt"
	"net"
)
//请求报文格式,一共有四部分组成
//第一个叫做 :请求行  第一行是请求方式/一种是get/另外一种叫做Post  （get一般都是明文的，post是加密的是看不到的 ）
//第二个叫做 :请求头
//第三个有个空行
//第四个 :包体（body）
func main() {
	//监听
	listener,err := net.Listen("tcp",":8000")
	if err != nil{
		fmt.Println(" net.Listen err",err)
		return
	}
	defer listener.Close()

	//阻塞等待用户的连接
	conn,err1 := listener.Accept()
	if err1 != nil{
		fmt.Println("listener.Accept err",err1)
		return
	}
	defer conn.Close()

    //接收客户端的数据
	buf := make([]byte,1024*4)
	n,err1 := conn.Read(buf)
	if n == 0{
		fmt.Println("conn.read err",err)
		return
	}
	fmt.Printf("#%v#",string(buf[:n]))
}
