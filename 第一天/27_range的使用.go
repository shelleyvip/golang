package main

import "fmt"

func main(){
	zhou := "milu"
	//迭代打印每个元素 默认返回两个值 一个是元素的位置，另一个是元素的本身
	for i,data := range zhou{
		fmt.Println("zhou[%d]=%c\n",i,data)
	}

}