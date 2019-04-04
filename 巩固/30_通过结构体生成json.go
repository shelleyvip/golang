package main

import (
	"encoding/json"
	"fmt"
)

//要想生成json格式成员首字母必须大写
type IT struct {
	Company string
	Subjects []string
	LsOk  bool
	Price 	float64
}
/*
type IT struct {
	Company string   `json:"-"`  代表此字段不会输出到屏幕
	Subjects []string  `json:"subjects"`  代表转化为小写
	LsOk  bool   	`json:"string"` 代表转换为字符串
	Price 	float64   `json:"string"` 代表转换为字符串
 */

func main()  {
	//定义一个结构体 同时初始化
      s := IT{"hello",[]string{"GO","Python","Test"},true,3.14}

      buf,err := json.MarshalIndent(s,""," ")
      if err != nil{
      	fmt.Println("err=",err)
		  return
	  }
      fmt.Println(	"buf=",string(buf))
}
