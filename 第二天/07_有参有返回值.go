package main

import "fmt"

func Shelley(a,b int)(zhou,xue int){
	if a > b {
		zhou = a
		xue = b
	}else{
		zhou = b
		xue = a
	}
	return //有返回值就必须有return
}
func main(){
	zhou,xue := Shelley(888,999)
	fmt.Printf("zhou=%d,xue=%d",zhou,xue)
	//通过匿名变量丢弃某个返回值
	//a 是元素本身_下标是丢弃的那个数用_下标丢掉
	a,_ := Shelley(888,999)
	fmt.Printf("a=%d\n",a)
	}
//通过匿名变量丢弃某个返回值
