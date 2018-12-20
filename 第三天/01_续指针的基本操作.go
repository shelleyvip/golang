package main

import "fmt"
//* 看到这个就是指针
//& 取地址操作
func main(){
	var a int = 90 //定义变量a赋值成90
	fmt.Println("a=",a)
	fmt.Printf("&a=%v\n" ,&a)

	///保存一个变量的地址，需要指针类型 *int 保存 int的地址，**int 保存 *int的地址
	//
	var p *int
	p = &a //指针指向谁，就把谁的 指针地址 赋值 给 指针变量
	fmt.Printf("p=%v\n,&a=%v\n",p,&a)
	*p = 666//*p 不是p的内存，是p所指向的内存（就是a）
	fmt.Printf("*p=%v\n &a=%v\n",*p,a)
	}