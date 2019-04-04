package main

import "fmt"

type PersonA struct {
	id int
	name string
	sex byte

}
type Student struct {
	PersonA
	age int
	addr string
}

func (tmp *PersonA)PrintlnInfoA()  {
	fmt.Printf("id=%d name =%s sex =%v",tmp.id,tmp.name,tmp.sex)
}

func (tmp *Student)PrintlnInfoA(){
	fmt.Println("tmp=",tmp)
}
func main() {
	s1 := Student{PersonA{666,"sheley",'m'},26,"bj"}
	//问题：这里的调用是调用的哪一个方法，答案是 func (tmp *Student)PrintlnInfoA(){
	//先着本作用域的方法，找不到在找继承的方法
	s1.PrintlnInfoA()

	//如果非要调用指定的某一个方法 那么就使用显示调用
	s1.PersonA.PrintlnInfoA()
}

