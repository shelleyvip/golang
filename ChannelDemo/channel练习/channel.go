package main

import "fmt"

type Cat struct {
	Name string
	Age int
}


func main() {
	//定义一个可以存放任何类型的管道 3个数据
	//var allChan interface{}
	allChan := make(chan interface{},3)

	allChan <- 10
	allChan <- "tom Jack"

	cat := Cat{"小花猫",3}
	allChan <- cat

	//我们希望获取到管道中的第三个元素，则先将前2个推出
	<- allChan
	<- allChan

	newCat := <- allChan
	fmt.Printf("newCat=%T newCat=%v",newCat,newCat)//看下第三个元素输出的是什么

	//以下的写法是错误的 编译器不通过
	//fmt.Printf("newCat.Name=%v",newCat.Name)

	//取出来的方法是用类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v",a.Name)
	}

