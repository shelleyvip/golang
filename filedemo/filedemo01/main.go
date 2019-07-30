package main

import (
	"fmt"
	"os"
)

func main() {
	//打开文件
	//概念说明
	//1.file 又叫 file对象
	//2.file 又叫 file指针
	//3.file 又叫 file句柄
	file,err := os.Open("/Users/golang/desktop/文本处理.png")
	if err != nil{
		fmt.Println("open file err",err)
	}
	//在这里输出文件看看文件是什么 看出file就是一个指针 *file
	fmt.Printf("file=%v",file)
	//关闭文件
	err = file.Close()
	if err != nil{
		fmt.Println("close file err",err)
	}

}
