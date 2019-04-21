package main

import "fmt"

type Student struct {
	Id int
	Name string
}


func main() {
	m := make(map[string]*Student)
	m["shelley"] = &Student{Id:1,Name:"周雪莉"}
	m["shelley"].Id = 2
   fmt.Println("m=",m)

	m1 :=make(map[string]*int)
	n := 1
	//m1["a"] = 1 err 不能直接修改
	m1["a"] = &n
	fmt.Println(n)
}
