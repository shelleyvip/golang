package main

import "fmt"

func main(){
	//切片和数组的区别
	//数组[]里面的长度是固定的常量，数组不能修改长度，len和cap永远是声明多少就是多少 是固定的
	a :=[]int{}
	fmt.Printf("len=%d,cap=%d\n",len(a),cap(a))

	b := []int{}
	fmt.Printf("len=%d,cap=%d",len(b),cap(b))
	//还可以用内置函数append给切片末尾增加一个成员
	b = append(b,11)//也就是说用append函数给len/cap的末尾增加1
	fmt.Printf("append:len=%d,cpa=%d",len(b),cap(b))
}
