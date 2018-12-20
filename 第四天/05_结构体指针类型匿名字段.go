package main

import "fmt"

type shelley struct {
	name string
	age int
	addr string
}

type zhou struct {
	*shelley  ///指针变量
	sex string
	id int
}
func main(){        ///取地址&Shelley 不加内容那么打印出来只是地址  //加内容猜是取它的内容
	s1 := zhou{&shelley{"周",26,"北京"},"女",888}
	fmt.Println(s1.name,s1.age,s1.addr,s1.sex,s1.id)
	//👆是第一种方式:直接初始化打印出来


	//这里是第二种给它先new一个新的空间来储存它  也就是先给你分配空间，在给你赋值，两种方式
    var s2 zhou
	s2.shelley = new(shelley)
	s2.name = "yoyo"
	s2.age = 99
	s2.addr = "朝阳"
	s2.sex = "男"
	s2.id = 888
	fmt.Println(s2.name,s2.age,s2.addr,s2.sex,s2.id)
}
