package main

import "fmt"

func main() {

	var shelley int
	fmt.Printf("请输入楼层：") //打印输出
	fmt.Scanln(&shelley) //scan函数接受用户输入

	switch shelley { //swich后面跟上 变量体本身
	case 1:
		fmt.Println("你输入的是1楼")

	case 2:
		fmt.Println("你输入的是2楼")
	case 3:
		fmt.Println("你输入的是3楼")
	case 4:
		fmt.Println("你输入的是4楼")

	default:
		fmt.Println("达不到你所想要的楼层，请跳楼吧")
	}
}