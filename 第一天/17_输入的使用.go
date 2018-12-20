package main

import "fmt"

func main(){
	var a int
	fmt.Println("请输入取变量a:")
	//"Scan", 阻塞等待用户的输入
	//"&",是取xxx的地址的作用
	fmt.Scan(&a)
}