package main

import (
	"fmt"
)

func main()  {
	type zhou int64 //给int64改别名 "zhou"
	var a zhou
    //等价于"a"是int64
   fmt.Printf("a type is %T\n",a)

	//给多个类型起别名
	type(
		xue byte
		li int64
	)
	var b xue = 'a'
	var c  li = 10
	fmt.Printf("b=%v,c=%v",b,c)

}