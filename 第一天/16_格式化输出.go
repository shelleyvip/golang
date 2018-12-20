package main

import "fmt"

func main(){
	a := 88
	b := 3.14
	c := "nihao"
	d := 'a'
	//%T操作变量所属类型
	fmt.Printf("a=%T,b=%T,c=%T,d=%T",a,b,c,d)
	//自动匹配格式输出
	fmt.Printf("a=%d,b=%f,c=%s,d=%c\n",a,b,c,d)
	//%v自动匹配格式输出
	fmt.Printf("a=%v,b=%v,c=%v,d=%v\n" ,a,b,c,d)
}

     //%d 整型
     //%f 浮点类型
     //%s 字符串类型
     //%c 字符类型