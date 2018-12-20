package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	list := os.Args
	if len(list) != 3{
		fmt.Println("使用方法：原文件 目的文件")
		return
	}

	srcFileName := list[1]
	desFileName := list[2]
	if srcFileName == desFileName{
		fmt.Println("两个值文件名字不可以相同")
		return
	}

	sf,err := os.Open(srcFileName)
	if err != nil{
		fmt.Println("err=",err)
		return
	}

	df,err := os.Create(desFileName)
	if err != nil{
		fmt.Println("err=",err)
		return
	}

	defer sf.Close()
	defer df.Close()



	buf := make([]byte,1024*4)
	for{
		n,err := sf.Read(buf)
		if err != nil{
			if err == io.EOF{
				break
			}
			fmt.Println("已完成拷贝")
		}
		df.Write(buf[:n])


	}
}