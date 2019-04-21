package main

import "fmt"

type Person struct {
	Name string
}

//给Person这个类型绑定一个方法  叫test
//test方法和Person类型绑定
//test方法只能通过Person类型的变量来调用，而不能直接调用，更不能使用其它变量来调用
func (p Person)test()  {
	fmt.Println("test()",p.Name)
}

func main() {
    var p Person
    p.Name = "tom"  //赋值
    p.test() //调用方法
}
