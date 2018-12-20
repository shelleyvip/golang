package main

import (
	"fmt"
	"os"
)
//获取文件属性 首先从路径build一个可执行文件 亮红后 使用./build的文件（注意:./后面不需要加.go）
func main()  {//后面加上你想要查询的文件名称的路径（必须是有效的文件）

	list := os.Args
	if len(list) != 2 { //usage：用法 使用
		fmt.Println("usage: xxx file")
		return
	}
	fileName := list[1]
	info,err := os.Stat(fileName) //获取文件函数的状态
	if err != nil {
		fmt.Println("没到找到文件错误",err)
		return
	}
       fmt.Println("文件的名称是->",info.Name())
       fmt.Println("文件的字节数->",info.Size())
       fmt.Println("文件的权限->",info.Mode())
}
