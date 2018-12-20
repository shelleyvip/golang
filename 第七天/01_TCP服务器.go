package main

import (
	"fmt"
	"net"
)
//使用方法：/先提供服务，在用 run 和build 都可以（适用于所有的程序）
//先运行tcp服务端，在运行客户端
func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listener.Close()

	//阻塞 等待用户连接 接受用户请求 Accept接受
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		//接收用户请求
		buf := make([]byte, 1024) //1024大小的缓冲区
		//n这个返回值就是它的长度 指定读了多少就打印多少
		n, err1 := conn.Read(buf)
		if err1 != nil {
			fmt.Println("conn.Read err1", err1)
			return
		}
		fmt.Println("buf=", string(buf[:n])) //:n指定读了多少就打印多少

		defer conn.Close() //关闭当前用户连接

}