package main

import "fmt"

func MaxAndMin(a,b int)(max,min int)  {
	if a > b{
		max = a
		min = b

	}else {
		max = b
		min = a

	}
	return  //只要有返回值的函数 必须return返回
	
}

func main() {
    max,min := MaxAndMin(99,88)
    fmt.Printf("max=%d,min=%d",max,min)

}
