package main

import "fmt"

func main(){
	a := 10
	b := "abc"
	f1 := func() {
		fmt.Println("a=",a)
		fmt.Println("a=",b)
	}
	f1()
}


//第二种
//func main(){
//	func(xue,li int){
//		fmt.Printf("xue=%d li=%d",xue,li)
//	}(000,999)
//}
