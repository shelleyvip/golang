package main

import (
	"fmt"
	"os"
)

func main(){
	//os.Stdout.Close()//这条指令是关闭后无法输出
	//fmt.Println("are you ok?")

   //标准设备文件（os.Stdont）已默认打开用户可以直接使用
	//os.Stdont
	os.Stdout.WriteString("are you ok?")
//WriteString 写入字符串

	//os.Stdin.Close()  //；这条指令是关闭后武法输入，直接是阻止输入
	var a int
	fmt.Println("请输入a:")
	fmt.Scan(&a)///从标准输入设备中读取内容放在a中
	fmt.Println("a=",a)
}
//Stdont 标准-输出
//Stdin  标准-输入
//Stderr 标准-错误


