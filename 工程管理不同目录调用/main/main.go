package main

import (
	"fmt"
	"golang/工程管理不同目录调用/model"
)

func main() {
	////创建一个student的实例
	//var stu = model.Student{Name:"tom",Score:78.9}
	//fmt.Println(stu)

	//当student结构体首字母小写 我们可以通过工厂模式来解决
	var stu = model.NewStudent("tom",88.8)
	fmt.Println(*stu)
	fmt.Println("name=",stu.Name,"score=",stu.Score)

	///如果Score字段首字母为小写，则，在其它宝不可以直接调用，我们可以提供一个方法
	fmt.Println("name=",stu.Name,"score=",stu.GetScore())

}

