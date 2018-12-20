package main

import "fmt"

type Shelley struct {
	name string
	age int
	sex string
}
//Shelley类型 实现一个方法
func (tmp *Shelley)Addr01(){
	fmt.Printf("name=%v age=%v sex=%v",tmp.name,tmp.age,tmp.sex)
}

//继承Shelley的字段 成员和方法都继承z了
type Zhou struct {
	Shelley  //匿名字段
	id int
	addr string
}
func main(){
	s := Zhou{Shelley{"周",26,"不指定类型"},128,"窝里"}
	s.Addr01()
}
