package main

import "fmt"
//	golang中的方法作用在指定的数据类型上的（即：和指定的数据类型绑定），因此自定义类型，
//都可以有方法，而不仅仅是struct 还可以是 int float32等都可以有方法


type integer int

func (i integer)Println()  {
    fmt.Println("i=",i)
}
//编写一个方法改变i的值
func (i *integer)change()  {
    *i = *i + 1
}

func main() {
	var i integer = 10
	i.Println()
	i.change()
	fmt.Println("i=",i)
}
