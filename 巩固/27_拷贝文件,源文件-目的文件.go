package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	list := os.Args //获取命令行参数

	if len(list) != 3{
		fmt.Println("usage:xxx srcFile dstFile")
		return
	}

	srcFileName := list[1]
	dstFileName := list[2]

	if srcFileName == dstFileName{
		fmt.Println("源文件和目的文件名字不能相同")
		return
	}
	//以只读方式打开源文件
	SF,err1 := os.Open(srcFileName)
	if err1 != nil{
		fmt.Println("err1=",err1)
		return
	}
    //新建目的文件
	dF,err2 := os.Create(dstFileName)
		if err2 != nil{
			fmt.Println("err2=",err2)
			return
		}
	//操作完毕需要关闭文件
	defer SF.Close()
	defer dF.Close()

	//核心文件处理 从源文件读组内容  往目的文件写 读多少写多少
	buf := make([]byte,4*1024)  //使用Mike 给它一个4K缓冲区大小
	for{
		n,err := SF.Read(buf) //读取缓冲区的内容
		if err != nil{
			if err == io.EOF{   //读取完毕
				break
				//fmt.Println("err=",err)
			}
			dF.Write(buf[:n])   //往目的文件写 读多少写多少
		}


	}





	}


