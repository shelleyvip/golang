package main


import "fmt"

type PersonA struct {
	id int
	name string
	sex byte
}

func (p PersonA)SetInfoValue()  {
	fmt.Printf("PrintlnValue = %p %v\n",&p,p)
}
func (p *PersonA)SetInfoPointer() {
	fmt.Printf("SetInfoPointe = %p %v\n", p, p)
}

func main() {
	s1 := PersonA{666,"mike",'m'}
	//fmt.Printf("main:s1=%p,%v",&s1,s1)

	f := (PersonA).SetInfoValue //方法表达式隐藏了接收者 显示把接收者传递过去
	f(s1)//显示把接收者传递过去 接收者s1  最终等价于调了这个函数SetInfoValue

	f2 := (*PersonA).SetInfoPointer
	f2(&s1)//显示把接收者传递过去 接收者s1  最终等价于调了这个函数SetInfoPointer
}
