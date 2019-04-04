package main

import "fmt"

type PersonA struct {
id int
name string
sex byte
}

func (p PersonA)SetInfoValue()  {
	fmt.Printf("PrintlnValue = %p %v",&p,p)
}
func (p *PersonA)SetInfoPointer() {
	fmt.Printf("SetInfoPointe = %p %v",p,p)
}
func main() {

	s1 := PersonA{666,"mike",'m'}
	fmt.Printf("main = %p %v",&s1,s1)
	//传统调用方式
	s1 .SetInfoPointer()

	//保存方式入口地址
	//好处是无需在调用函数 指定米一个方法
	pFunc := s1.SetInfoPointer //这个就是方法值 调用函数时无需在传递接收者，隐藏了接收者
	pFunc()    //这个就等价于SetInfoPointer（）

	vFunc := s1.SetInfoValue
	vFunc() //这个就等价于	SetInfoValue() 的方法



}