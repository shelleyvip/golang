package main

import (
	zhou "fmt" //给包起一个别名,在包前面加别称
	//_ "fmt"
	///_其实就是引入该包 而不使用包里的函数，而是调用了该包里面的init函数
)
func main(){
	zhou.Println("nihao")//相当于fmt
}