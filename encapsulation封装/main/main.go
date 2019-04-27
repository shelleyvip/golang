package main

import (
	"fmt"
	"golang/encapsulation封装/model"
)

func main() {
	p := model.NewPerson("smith")
	p.Setage(18)
	p.Setsal(5000)
	fmt.Println(p)
	fmt.Println(p.Name,"age=",p.Getage(),"sal=",p.Getsal())
}
