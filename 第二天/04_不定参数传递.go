package main

import "fmt"

func Mynan(milu ...int){
	//第一个是下标 第二个是下标对应的数
	for _,data := range milu{
		fmt.Println("data=",data) ////用不定参数的时候可以通过迭代 for _,data :=range xxx{
	}
}

//func Myhu(zailing ...int){
	//Mynan(zailing...)     ///<全部元素传递

//}


func Myhu(zailing ...int){
	Mynan(zailing[2:]... )////[2:]是不包括2本身传递
	Mynan(zailing[:2]... )/// [:2]是包括本身传递的 左避右开的原则
}

func main(){
	Myhu(1,2,3,4,5,6,7)
}

