package main

import (
	json2 "encoding/json"
	"fmt"
)

type Monster struct {
	Name string`json:"name"`
	Id int`json:"id"`
	Age int`json:"age"`
}

func main() {
	//创建一个monster变量
   monster := Monster{"牛魔王",666,500}
   jsonStr,err := json2.Marshal(monster)
   if err != nil{
   	fmt.Println("json 处理错误",err)
   }
   fmt.Println("jsonStr",string(jsonStr))
}
