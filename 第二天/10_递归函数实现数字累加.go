package main

import "fmt"

//func Shelley01 ()(zhou int){
	//if i := 1; i <= 100; i++ {
	//	zhou += i
	//}
	//return
//}

func zailing (a int)int{
	if a == 1{
		return 1
	}
	return a + zailing(a-1)
}

func main(){
	zhou := zailing(100)
	fmt.Println("zhou=",zhou)
}