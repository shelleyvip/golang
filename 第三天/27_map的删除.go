package main

import "fmt"

func main(){
	s := map[int]string{1:"big",2:"baby",3:"boy"}
	fmt.Println("s=",s)
	delete(s,3)//delete 是删除某以为指定的元素
	fmt.Println("s=",s)
}
