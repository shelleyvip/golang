package main

import "fmt"

type Shelley struct {
	name string
	age int
	sex string
}
//带有接受者的函数叫方法
func(tmp Shelley)value(){
	fmt.Println("tmp=",tmp)

}

//通过一个函数给成员赋值
func (tmp *Shelley)rsult(a string ,b int,c string){
	tmp.name = a
	tmp.age = b
	tmp.sex = c
}

func main(){
	//定义同时初始化
	s1:=Shelley{"nn",88,"n"}
	s1.value()

	//定义一个结构体变量
	var s2 Shelley
	(&s2).rsult("周",25,"女")
	s2.value()

}