package main

import "fmt"

type Student struct {
	Name string
	Id int
}


func main() {
	s1 := Student{"xiaoxiao",1}
	fmt.Println(s1)
}
