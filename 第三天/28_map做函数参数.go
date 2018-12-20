package main

import "fmt"
func Shelley(a map[int]string){
	delete(a,2)  //用delete删除指定的某一个key键
}

func main(){
	a := map[int]string{1:"周",2:"雪",3:"莉",4:"你好"}
	Shelley(a)  //调用Shelley下的内容 并执行
	fmt.Println("a=",a)
}