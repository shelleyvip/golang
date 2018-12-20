package main

import "fmt"

type Shelley struct {
	name string
	age int
	sex string
}

func (tmp Shelley)Addr(){
	fmt.Println("Addr")
}

func (tmp *Shelley)Pointer(){
	fmt.Println("Pointer")
}

func main(){
	//结构体变量是一个指针变量，它能够调用哪些方法，这些方法就是一个集合，简称方法集
	p1 := &Shelley{"周",88,"狼性呵呵"}
	//p1.Pointer()//这个调用的是 func (tmp *Shelley)Pointer()
	(*p1).Pointer()//把*P转为P在调用等价于👆上面

	p1.Addr()///这个为什么能调用Addr呢？
	//因为它 内部做了转换先把 指针P1 转换成*p1  在调用
	//(*p1).Addr()

}
