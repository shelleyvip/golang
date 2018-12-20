package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
//每次读取一行
func ReadFileLine(path string)  {
	//打开文件
	f,err := os.Open(path)
	if  err !=nil  {
		fmt.Println("err",err)
		return
	}
	//关闭文件
	defer f.Close()
	//新建一个缓冲区，把内容先放在缓冲区
	r := bufio.NewReader(f)  //也就是说把f内容放进来
	for {             //遇到'\n'结束读取,但是同时也把\n读取进去
		buf,err := r.ReadBytes('\n')
		if err != nil{
			if err == io.EOF{
				break
			}
			fmt.Println("err=",err)
		}
		fmt.Printf("buf=#%s#\n",string(buf))
	}
}

func main()  {
	//path := "./读文件.txt" //相对路径读取
	path := "/Users/golang/Work/读文件.txt" ////绝对路径读取

	ReadFileLine(path)//读文件
	//reed：读
	// file：写
	//line：行
}



