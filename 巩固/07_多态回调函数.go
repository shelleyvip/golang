package main

import "fmt"

type FuncType func(int,int)int

func add(a,b int)int  {
	return a +b
}
func mul(a,b int) int {
	return a * b

}

func minus(a,b int)int  {
	return a -b
	
}

func Calc(a,b int,fTest FuncType)(result int){
	//fmt.Println("Calc")
	result = fTest(a,b)
	return
}
func main()  {
   a := Calc(10,20,add)
   b := Calc(20,10,minus)
   c := Calc(88,99,mul)

   fmt.Println("c=",c)
   fmt.Println("b=",b)
   fmt.Println("a=",a)
}
