package main

import "fmt"

func main(){
	var a int = 10
	fmt.Println("a=",a)//每个变量有两层含义：1变量的内存 2变量的地址
	fmt.Printf("&a=%v",&a)//打印a的地址用%v 或者%p
	///"&"是取a的地址
	///a是内存的内容  &a是外面的符号
}