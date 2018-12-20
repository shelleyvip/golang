package main

import "fmt"
//定义一个结构体 用type 和struct
type Shelley struct {
	name string
	id int
	age int    //这里是定义类型 下面.操作 写相应的内容
	addr string
	sex string
}
func main(){
	//定义一个结构体普通变量
	var s Shelley
	//操作成员需要使用.操作运算符
	s. name = "周"
	s. id = 123
	s. age = 25
	s. addr = "北京"
	s. sex = "nv"
	fmt.Println("s=",s)

}
