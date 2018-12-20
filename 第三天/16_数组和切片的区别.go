package main

import "fmt"

func main()  {

//切片和数组的区别
//数组【】里面的长度是一个固定的常量,数组不能修改len 和cap 是多少就是多少
	a :=[5]int{1,2,3,4,5}
	fmt.Printf("len=%d cap=%d",len(a),cap(a))


//切片里面为空或是...切片的长度和容量可以不固定
b := [...]int{1,2,3,4,5,6}
fmt.Printf("b的长度是%d\nb的容量是%d\n",len(b),cap(b))


//借助make函数，格式make（切片类型，长度，容量）
c := make([]int,5,10)  //第一个是长度第二个是容量
fmt.Printf("c=len长度 %v\nc=cap容量 %v\n",len(c),cap(c))


d := make([]int,10)//若只有一个那么len 和cap 都是10
fmt.Printf(" d=长度%d\nd=的容量%d\n",len(d),cap(d))

}