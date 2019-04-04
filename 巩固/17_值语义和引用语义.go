package main

import "fmt"

type Preson struct {
	id int
	name string
	age int
}
//修改成员变量的值
//接收者为普通变量 非指针 也就是说是 值语义 拷贝了一份
func (p Preson)SetInfoValues(a int,b string,c int)  {
      p.id = a
      p.name = b
      p.age = c
      fmt.Println("p=",p)
      fmt.Printf("SetInfoValues &p p=%p\n",&p)
}
//这个接收者为指针变量 引用类型
func (p *Preson)SetInfoPointer(a int, b string,c int)  {
	p.id = a
	p.name = b
	p.age = c
	fmt.Println("p=",p)
	fmt.Printf("SetInfoValues &p p=%p\n",p)
}
//func (tmp person)Println()  {
//	fmt.Println("tmp=",tmp)
//}
func main() {
	p1 := Preson{1,"米露",3}
     fmt.Printf("&p1 = %p\n",p1)//打印地址
	 //
     //p1.SetInfoValues(666,"大胖妹",3) //值语义
     //fmt.Println("p1=",p1)

	//引用语义
	(&p1).SetInfoValues(777,"mmm",666)
	fmt.Println("p1=",p1)




}
