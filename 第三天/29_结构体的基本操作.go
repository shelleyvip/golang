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
	                 ///相对等的类型
	zhou := Shelley{name:"周雪莉",age:22,sex:"中",qq:128,addr:"朝阳"}
	fmt.Println("zhou=",zhou)

}
