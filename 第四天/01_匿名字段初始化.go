package main

import "fmt"

type shelley struct{
	name string
	age int
	addr string
}
type zhou struct{
	shelley//只有类型，没有名字，也就是说继承了Shelley的成员
	id int // 这里就是 zhou 也有自己的成员
	sex string //这里就是 zhou 也有自己的成员
}
func main(){//shelley 也是一个结构体所以一定要加{}否则会err的
	li := zhou{shelley{"周雪莉",25,"北京"},128,"女"}
	fmt.Println("li=",li)//第一个Shelley{是Shelley的成员} //后面追加的是zhou的成员
	//格式化打印+v打印显示更详细
	fmt.Printf("li=%+v\n",li)

	//指定成员初始化，没有初始化的自动赋值为0
	li1 := zhou{sex:"男"}
	fmt.Printf("li1=%+v\n",li1)
	/////shelley 也是一个结构体所以一定要加{}否则会err的
	///两个结构体 给指定的某个成员初始化
	li2 := zhou{shelley:shelley {name:"周"},id:1}
	fmt.Printf("li2=%+v",li2)

}

