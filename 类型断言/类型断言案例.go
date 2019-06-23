package main

import "fmt"

type Point struct {
	x int
	y int
}
func main()  {
	var a interface{}
	var point Point = Point{1,2}
	a = point  //这样转换没问题

	//如何将a赋值给point变量
	var b Point
	//b = a //不可以
	//b = a.(point)可以
     b = a.(Point)  //b=a.(point)就是类型断言，表达判断a是否指向point类型的变量，如果是就转成point
     fmt.Println(b) //类型并赋值给b变量，否则报错

     //类型断言其它案例
     var x interface{}
     var b2 = float32(1.1)
     x = b2
     y := x.(float32)
     fmt.Printf("y的类型是%T 值是%v",y,y)

     //带检测机制的断言
     y,ok := x.(float32)
     if ok{
     	fmt.Println("convert success")
     	fmt.Printf("y的类型是%T 值是%v",y,y)
	 }else {
	 	fmt.Println("convert fail")

	 }
     fmt.Println("继续执行" )
}
