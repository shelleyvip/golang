package main

import "fmt"
//modify a 指向实参数组a 他是指向数组 所以是数组指针
//*a代表所指向的内存，就是实参a
func modify(a *[5]int)  {
	(*a)[3]= 666    //
	fmt.Println("*a=",*a)

}


func main() {
	a :=[5]int{1,2,3,4,5}  //数组初始化
	fmt.Println("main a=",a)
	modify(&a)  //地址传递
}
