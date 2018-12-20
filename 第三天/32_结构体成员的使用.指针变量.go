package main

import "fmt"
//定义以个结构体 用type 和struct
type Shelley struct {
	name string
	age int
	sex string    ///定义类型
	qq int
	addr string
}

func main(){
	//// 第一种指针变量
	var s Shelley
	var p1 *Shelley
	p1 = &s
	(*p1).name = "祖宗"
	(*p1).age = 38
	(*p1).sex = "男"    //若是单引号对应的就是assc码
	(*p1).qq = 128
	(*p1).addr = "bj"
	fmt.Println("s=",s)


     ///第二种指针变量的方法 new一个新的空间
	p2 := new(Shelley)
	(*p2).name = "宗"
	(*p2).age = 38
	(*p2).sex = "男"    //若是单引号对应的就是assc码
	(*p2).qq = 128
	(*p2).addr = "bj"
	fmt.Println("p2=",p2)

}

