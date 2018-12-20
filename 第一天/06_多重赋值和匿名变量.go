package main

import "fmt"

func main(){
	//多重赋值
	a,b := 10,20
	a,b = b,a
	fmt.Printf("a=%d b=%d\n",a,b)
	//匿名变量
	var zhou int  //给变量取一个名字叫zhou
	i := 88       // 然后声明两个变量
	j := 99
	zhou,_=i,j    //如果想打印其中一个的变量值 用_下标 作用是：匿名变量 丢弃后面的那个数据不处理
	fmt.Println("zhou=",zhou)



}
