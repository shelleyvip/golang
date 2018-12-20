package main

import "fmt"

type setinfo interface {//这个是子集
	object()
}
type Student interface {//这个是超集 因为它范围比较多的
	setinfo  //匿名字段，继承了object（）的接口
	result(zhou string)//	在这里我自己student还有一个方法，这个方法带有一个参数
}
type Shelley struct{
	name string
	id int
}

func (tmp *Shelley)object(){
	fmt.Printf("Shelley[%s %d]\n",tmp.name,tmp.id)
}
func(tmp *Shelley)result(zhou string){
	fmt.Println("Shelley在唱歌:",zhou)
}
func main() {
	//定义一个接口类型的变量    /就定义超集/因为它范围比较多
	var i Student
	s := &Shelley{"make", 666}
	i = s
	i.object()  //继承过来的方法
	i.result("学生歌")//
}