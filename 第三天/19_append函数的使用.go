package main

import "fmt"

func main(){
	s1 :=[]int{}
	fmt.Println("s1=",s1)
	fmt.Printf("len=%d,cap=%d\n",len(s1),cap(s1))
	//append是内置添加函数
	s1 = append(s1,88)

	//在原切片末尾添加元素
	s2 :=[]int{1,2,3}
	s2 = append(s2,5)
	s2 = append(s2,5) //添加后的函数会在123后面出现
	s2 = append(s2,5)
	fmt.Println("s2=",s2)
	fmt.Printf("len=%d,cap=%d\n",len(s2),cap(s2))

}
