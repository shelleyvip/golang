package main

import "fmt"

func main()  {
	const(
		a = iota////iota是给常量赋值使用的
		b = iota ////是常量的自动生成器，每隔一行 自动累加1
		c = iota
	)
	fmt.Printf("a=%d,b=%d,c=%d",a,b,c)




const e = iota   ////iota 遇到const 重置为零
fmt.Printf("e=%d",e)

const(
	h = iota
	j,f,w = iota,iota,iota/////如果是在同一行 它的值都一样
	r = iota
)
fmt.Printf("h=%d,j=%d,f=%d,w=%d,r=%d\n",h,j,f,w,r)







}