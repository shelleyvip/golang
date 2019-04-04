package main

import "fmt"

func MyDiv(a,b int)(result int,err error)  {
	err = nil
	if b == 0{
		fmt.Println("分母不能为0",err)
	}else {
		result = a/b
	}
	return
	
}
func main() {
	result,err := MyDiv(20,10)
	if err != nil{
		fmt.Println("err=",err)
	}else {
		fmt.Println("result=",result)
	}
}
