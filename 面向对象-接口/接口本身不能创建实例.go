package main

import "fmt"

type Ainterface interface {
	Say() //里买又一个Say的接口
}
type Stu struct {
	Name string
}

func (stu Stu)Say()  {
    fmt.Println("Stu Say()")
}

//只要是自定义类型 就可以实现接口 不仅仅是结构体类型
type integer int //普通的自定义类型 也可以实现接口 只要把它的方法实现了 就实现这个接口了

type Age struct {
	Sex string
}

func (a Age)Say()  {
	fmt.Println("这里是另外一个例子")
}

func (i integer)Say()  {
	fmt.Println("integer Say i=",i)
}

func main() {
	var stu Stu //结构体变量实现了Say实现了A
	//接口本身不能创建实例，但是可以指向一个实现了该接口的自定义变量（实例）
	var a Ainterface = stu
	a.Say()

	var i integer = 10 //创建一个接口变量
	var b Ainterface = i
	b.Say()

	var c Age  //把age赋给c
	var d Ainterface = c //
	d.Say()   //指向该接口
}

