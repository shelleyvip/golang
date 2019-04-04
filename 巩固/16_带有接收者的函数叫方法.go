package main

import "fmt"

type Person struct {
	id int
	name string
	age int
	addr string
}
//带有接收者的函数叫方法 PrintlnInfo ()是接收者 它 接收(tmp Person)的方法
//接收者PrintlnInfo 接收了(tmp Person）结构体 并实现打印的方法
func (tmp Person)PrintlnInfo () {
	fmt.Println("tmp=",tmp)
	//fmt.Printf("tmp=%+v",tmp)
}

//在通过一个函数给成员赋值
//p 就是变量名它叫接收者 (p *person) 这个就叫做接收类型接收者
//(p *person）这一块只要类型不一样 即使是重名也是两个不同的方法
func (p *Person)SetInfo(a int,b string,c int,d string)  {
	p.id = a
	p.name = b
	p.age = c
	p.addr = d

}
func main()  {
	//定义同时初始化
	p := Person{1,"mike",25,"bj"}
	p.PrintlnInfo() //这边就使用PrintlnInfo的方法去打印

	//定义一个结构体变量
	var p2 Person
	(&p2).SetInfo(2,"sdsdsd",22,"sdsda")
	p2.PrintlnInfo()//在使用这使用PrintlnInfo的方法去打印
}
