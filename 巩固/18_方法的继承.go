package main

import "fmt"

type Person01 struct {
	id int
	name string
	age int

}
 //把person的成员都继承了
type Student struct {
	Person01
	addr string
	sex byte
}

func (tmp *Person01)PrintlnInfo()  {
	fmt.Printf("id=%d name= %s age=%d",tmp.id,tmp.name,tmp.age)
}

func main()  {
	s1 := Student{Person01{666,"shelley",26},"bj",'m'}
	s1.PrintlnInfo() //调用PrintlnInfo打印方法

}


