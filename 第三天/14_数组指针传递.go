package main

import "fmt"
func shelley(p *[5]int){

	///也就是说我把我的地址给到你之后，可以让你改变我的值

	//p指向实参数组a，由于p它是指向数组，所以它就是数组指针

	//*p代表指针所指向的内存，所以他就是是实参a
	(*p)[3]=88 //给指定的第几个元素赋值
	fmt.Println("*p=",*p)
}

func main(){
	a :=[5]int{1,2,3,4,5}
	shelley(&a) //这个&a 的地址是被拷贝了一份进shelley的
	             ///也就是说我把我的地址给到你之后，可以让你改变我的值
	fmt.Println("a=",a)
}