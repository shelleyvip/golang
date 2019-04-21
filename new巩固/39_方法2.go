package main

import "fmt"

//1)声明一个结构体Circle  字段为radius
type Circle struct {
	radius float64
}
//2)声明一个叫area的方法和Circle绑定，可以返回面积
func (c Circle)area() float64 {
	return 3.14 * c.radius * c.radius
}
func main() {
	//创建一个结构体变量
	var c Circle
	//给radius赋值
	c.radius = 4.0
	//接收者res c去调用这个方法 去实现方法内的功能
	res := c.area()
	fmt.Println("面积是:",res)

}