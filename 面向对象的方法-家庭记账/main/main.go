package main

import (
	"fmt"
	"golang/面向对象的方法-家庭记账/utils"
)

func main() {
   fmt.Println("这个是面向对象的方式完成的～～～")
   //调的是县创建这个结构体的指针实例    在调用的MainMenu()  在调用了这个结构体的各个方法
   utils.NewFamilyAccount().MainMenu()
}
