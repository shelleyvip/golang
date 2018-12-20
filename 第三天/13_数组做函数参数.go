package main

import "fmt"
//数组做函数参数它只是值传递 给shelley拷贝了一份而已 //所以（main包下面a的内容不会改变）
func shelley(a[5] int){
	a[3] = 666     //是参数组的(也就是main包下面的)每个元素 给形参数组拷贝一份
	//然后给指定的某一个元素赋值
	//元素从下标0开始算起，也就是说第一个数字是0 开始算起 也就是：1，2，3，666，5
	fmt.Println("a=",a)
}

func main(){
	a := [5]int{1,2,3,4,5}// 数组的初始化
	shelley(a)            //把数组传过去
	fmt.Println("a=",a)
}