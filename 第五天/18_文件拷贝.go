package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	list := os.Args //获取命令行参数
	if len(list) != 3 {
		fmt.Println("使用方法： 源文件srcFile 目标文件desFile")
		return
	}
	srcFileName :=list[1]
	desFileName :=list[2]
	if  srcFileName == desFileName {
		fmt.Println("两个值名称不能相同")
		return
	}
	//只读方式打开原文件
	sF,err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("err=",err)
		return
	}
  //新建文件
	dF,err1 := os.Create(desFileName)
	if err1 != nil {
		fmt.Println("err1=",err1)
		return
	}
	//操作完毕需要关闭
	defer  sF.Close()
	defer  dF.Close()
	//核心处理
	buf := make([]byte,1024*4)
	for {
		n,err := sF.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		fmt.Println("拷贝完成✅,请继续操作...")
		dF.Write(buf[:n])
	}

}
