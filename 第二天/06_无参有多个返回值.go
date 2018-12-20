package main

import "fmt"
//有多个返回值
func Shelley()(a,b,c int){  ///go 语言推荐写法
	a,b,c = 111,222,333
	return
}
func main(){
	a,b,c := Shelley() //调用Shelley01（）并打印出 a,b,c 各自相应的值
	fmt.Printf("a=%d b=%d c=%d\n",a,b,c)
}
