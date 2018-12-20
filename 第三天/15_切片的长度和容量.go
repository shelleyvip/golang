package main

import "fmt"

func main(){
	a :=[]int{1,2,3,4,5}
	s := a[1:4:5]  //从下标开始取 第一个是下标 第二个长度 第三个是容量
	fmt.Println("s=",s)
	fmt.Println("len(s)=",len(s)) //长度4-1 是3
	fmt.Println("cap(s)=",cap(s))  //容量 5-1 是4
}