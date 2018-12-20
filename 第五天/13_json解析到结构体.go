package main

import (
	"encoding/json"
	"fmt"
)

type IT3 struct {
	Company string `json:"company"`//那么结构体首字母必须为大写，这样下面又都是小写
	Subjects []string  `json:"subjects"`   //，那么在这里就只能通过字段来给它改成小写
	IsOk bool      `json:",isOk"` //所有的名字都来个二次编码
	Price float64 `json:",price"`
}
func main(){
	jsonBuf :=`{
 "Company": "itcast",
 "Subjects": [
  "Go",
  "C++",
  "Python",
  "Test"
 ],
 "IsOk": true,
 "Price": 666.666
}`

	var tmp IT3 //定义一个结构体变量
	err := json.Unmarshal([]byte(jsonBuf),&tmp)//就把结构体的内容解析到对应的字段
	if err != nil{
		fmt.Println("err=",err)
		return
	}else{
		fmt.Printf("tmp=%+v\n",tmp)
	}

	//在这里那么我只想要一个字段，那么就定义结构体的时候只有一个字段
	type IT4 struct {
		Subjects[]string `json:"subjects"`//二次编码
	}
	var tmp4  IT4
	err = json.Unmarshal([]byte(jsonBuf),&tmp4)
	if err != nil{
		fmt.Println("err=",err)
		return
	}else {
		fmt.Printf("tmp4=%+v\n",tmp4)
	}
}