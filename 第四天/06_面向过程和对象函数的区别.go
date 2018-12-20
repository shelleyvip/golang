package main

import "fmt"
//面向对象   方法：给某个类型绑定一个函数
type long int

func (temp long)Add01(zhou long)long{
	return temp + zhou
}
func main(){
	//定义一个变量
	var a long = 10
	//调用方法格式：变量名函数（所需参数）
	r := a.Add01(20)
	fmt.Println("r=",r)
}