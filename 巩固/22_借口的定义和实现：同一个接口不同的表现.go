package main

import "fmt"

type Humaner interface {
	sayhi()
}

type Student struct {
	id int
	name string
}

type Teacher struct {
	addr string
}

func (tmp *Student ) sayhi() {
  fmt.Printf("Student[%d %s]\n",tmp.id,tmp.name)
}
func (tmp *Teacher) sayhi() {
	fmt.Printf("Teacher= %s\n",tmp.addr)
}

type Mystr string

func (tmp *Mystr) sayhi() {
	fmt.Printf("Myst=%s\n",*tmp)

}

func main() {
	//定义接口类型变量
	var i Humaner
	fmt.Println(i)

	//只是实现了此接口方法的类型，那么这个类型的变量（接收者类型）就可以给i赋值
	s := &Student{666,"haha"}
	i = s
	s.sayhi() ////掉的接口都是一样，但是最终的表现却不一样

	s1 := &Teacher{"beijing"}
	i = s1
	s1.sayhi()//掉的接口都是一样，但是最终的表现却不一样

	var s2 Mystr = "hello world"
	i = &s2
	s2.sayhi()//掉的接口都是一样，但是最终的表现却不一样



}
