package main

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(path string)  {
	//打开文件
	f,err := os.Open(path)
	if  err !=nil  {
		fmt.Println("err",err)
		return

	}
	//关闭文件
	defer  f.Close()
	buf := make([]byte,1024*2) //2KB
	//从文件读取内容长度
	n,err1 := f.Read(buf)           //in EOF，这个代表结束符
	if err1 != nil && err1 != io.EOF {//文件出错，同时还没有结尾
		fmt.Println("err1=",err1)
		return

	}
	fmt.Println("buf =",string(buf[:n]))

	}

func main()  {
	//path := "./读文件.txt" //相对路径读取
	path := "/Users/golang/Work/读文件.txt"////绝对路径读取
	ReadFile(path)//读文件
}

