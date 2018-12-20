package main

import "fmt"
//类似...int类型 这样的类型 。...叫做不定参数
func zhou(nan ...int) {
	fmt.Printf("nan=%d\n", nan)
	//"len"是一个内置函数 是用来测你这个元素个数是多少位的
	fmt.Println("len(nan)=", len(nan))
}


func main(){
	zhou(1)
	zhou(1,2)///你可以传一个或多个参数
	zhou(2,3,4,5,6,7,8,9,0)
}
