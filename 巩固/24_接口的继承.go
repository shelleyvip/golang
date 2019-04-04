package main

import "fmt"

type Humaner0 interface {
	sayhi1()
}

type Personer1 interface {
	Humaner()
	sing(lrc string)
}

type Student0 struct {
	name string
	id int
}
//Student0 实现了sayhi1()
func (tmp *Student0)sayhi1() {
	fmt.Printf("Student0=%s %d",tmp.name,tmp.id)
}
//Student0有实现了lrc string 这个参数
func (tmp *Student0)sayhi1(lrc string)  {
  fmt.Println("Student0在唱着:",lrc)
}

func main() {
	//定义一个接口类型的变量
	var b Personer1
	a := &Student0{"ssss",333}
	b = a



}
