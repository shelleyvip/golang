package main

import "fmt"

func main(){
	a := 66
	//Println 是一段一段的打印
	fmt.Println("a=",a)
	//格式化输出把a的内容放在%d的位置
	//"a=10\n"这个字符串输出到屏幕；\n代表换行符
	fmt.Printf("a=%d\n",a)

	b := 90
	c := 100
	fmt.Printf("b=%d,c=%d\n",b,c)
}
///Println 是一段一段处理自动加换行 * 每一段要以逗号（,）分隔 //每一段要加上双引号""
 //Printf 是格式化输出把a的内容放在%d的上面 \n 是换行符 //无须每段都加双引号""
