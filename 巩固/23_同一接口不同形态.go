package main

import "fmt"

type Humaner interface {
	sayhi()  //定义接口
}

type Student struct {
	id int
	name string  //结构体
}

type Teacher struct {
	addr string   //结构体
}

type Mystr string  //自定义string类型

func (tmp *Student ) sayhi() {
	fmt.Printf("Student[%d %s]\n",tmp.id,tmp.name)
}
func (tmp *Teacher) sayhi() {
	fmt.Printf("Teacher= %s\n",tmp.addr)
}

func (tmp *Mystr) sayhi() {
	fmt.Printf("Myst=%s\n",*tmp)

}
func WhoSahi(i Humaner)  {
	i.sayhi()               //把接口封装到一个函数
	
}

func main() {
	s := &Student{6,"nihao"}
	t := &Teacher{"bj"}
	var str Mystr = "hello world"

	WhoSahi(s)
	WhoSahi(t)
	WhoSahi(&str)

	//还可以用切片来实现多态
	x := make([]Humaner,3)
	x[0] = s    //对应类型的参数
	x[1] = t
	x[2] = &str
	for _,i :=range x{
			i.sayhi() //把sayhi传给i
	}


}

//func main01() {
//
//	//定义接口类型变量
//	var i Humaner
//	fmt.Println(i)
//
//	//只是实现了此接口方法的类型，那么这个类型的变量（接收者类型）就可以给i赋值
//	s := &Student{666,"haha"}
//	i = s
//	s.sayhi() //掉的接口都是一样，但是最终的表现却不一样
//
//	s1 := &Teacher{"beijing"}
//	i = s1
//	s1.sayhi()  //掉的接口都是一样，但是最终的表现却不一样
//
//	var s2 Mystr = "hello world"
//	i = &s2
//	s2.sayhi()  //掉的接口都是一样，但是最终的表现却不一样
//
//
//
//}
