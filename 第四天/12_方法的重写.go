package main

import "fmt"

type Shelley struct {
	name string
	age int
	sex string
}
func (tmp Shelley)Addr01(){
	fmt.Printf("Shelley:name=%v age=%v sex=%v",tmp.name,tmp.age,tmp.sex)
}
//
type student struct {
	Shelley
	id int
	addr string
}
func (tmp student)Addr01(){
	fmt.Println("student:tmp=",tmp)
}

func main(){
	s1 := student{Shelley{"周家大院",25,"女"},123333,"北京"}
	//有两个同名的方法 那么到底调用的是Shelley的，还是student的 结论是student的
	s1.Addr01()//就近原则，先找本作用域的方法，找不到在用继承方法


	//如果非要把Shelley的方法调出来
	//就是使用显示调用/那么就指定要shelley里面的方法
	s1.Shelley.Addr01()
}