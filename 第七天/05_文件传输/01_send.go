package main

import (
	"fmt"
	"io"
	"net"
	"os"
)
//使用方法：build 或者run
//先编译成两个可执行文件，直接./运行（先go run recv等待接收）在运行./send(后面不需要.go 直接加入绝对路径既可以了)
func main()  {                                     //后面跟上绝对路径的有效文件
	sender()
}

	func sender() {
		//读取命令行，返回的是一个字符串切片，list[0] = "go run xxx.go" list[1] = "要发送的文件的绝对路径“
		list := os.Args //Args传感器传送的意思
		filepath := string(list[1]) //文件绝对路径转成字符串
		fileInfo, err := os.Stat(filepath) //通过stat函数得到文件属性（Name，Size ...）
		if err != nil {
			fmt.Println("os.Stat err:", err)
			return
		}
		fileName := fileInfo.Name() //获得文件名
		conn, err := net.Dial("tcp", "127.0.0.1:8008") //与接收端建立TCP连接
		if err != nil {
			fmt.Println("net.Dial err:", err)
			return
		}
		defer conn.Close() //不要忘记关闭
		conn.Write([]byte(fileName)) //把文件名发送给接收端
		buf := make([]byte, 16)
		n, err := conn.Read(buf) //接收"ok"信号
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		if string(buf[:n]) == "ok" { //如果是"ok",则发送文件内容
			sendFile(filepath, conn) //将文件的绝对路径传入，不是文件名
		}
	}
	func sendFile(filepath string, conn net.Conn) {
		fl, err := os.Open(filepath) //打开文件
		if err != nil {
			fmt.Println("os.Open err:", err)
			return
		}
		defer fl.Close()
		buf := make([]byte, 4096)
		for {
			n, err := fl.Read(buf) //读取到缓冲区
			if err != nil {
				if err == io.EOF { //判断是否读完
					fmt.Println("send file complete发送完毕✅")
				} else {
					fmt.Println("file read err:", err)
				}
				return
			}
			conn.Write(buf[:n]) //发送文件内容给接收端
		}
	}


