package main

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(path string)  {
	f,err := os.Open(path)
	if err != nil{
		fmt.Println("err=",err)
		return
	}
	defer f.Close()
	buf := make([]byte,1024*2)
	n,err1 := f.Read(buf)
	if err1 != nil && err1 == io.EOF{  //同时文件没有到结尾   io.EOF结束符
		fmt.Println("err1=",err1)
		return
	}
	fmt.Println("buf=",string(buf[:n]))
}



//写入文件
func WriteFile(path string)  {
	//打开新建文件
	f,err := os.Create(path)
	if err !=nil{
		fmt.Println("err=",err)
		return
	}
	//使用完毕需要关闭文件
	defer f.Close()

	var buf string
	for i := 0;i<10;i++{
		//fmt.Sprintf 这个的意思是把"i=%d\n"存储在buf中
		buf = fmt.Sprintf("i=%d\n",i)
		fmt.Println("buf=",buf)

		n,err := f.WriteString(buf)
		if err != nil{
			fmt.Println("err=",err)
		}
		fmt.Println("n=",n)
	}

}
func main() {
	path := "./demo.txt" //创建一个文本文件
	WriteFile(path)      //把这个文件传进来
}
