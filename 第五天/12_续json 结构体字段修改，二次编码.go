package main

import (
	"encoding/json"
	"fmt"
)

type IT2 struct {
	Company string `json:"company"`//这样Company就变成了小写   这个就属于二次编码
	Subjects []string  `json:"-"`//这种就代表此字段里面的内容不会输出到屏幕
	IsOk bool      `json:",string"`//这个就代表把这个bool类型转换为字符串在输出屏幕
	Price float64 `json:",string"`//这个也就以字符串的形式输出到屏幕
}

func main(){
	s1 := IT2{"istest",[]string{"Go","C++","Python","Test"},true,88.88}
	student,err := json.MarshalIndent(s1,""," ")
	if err != nil{
		fmt.Println("err=",err)
		return
	}else {
		fmt.Println("student=",string(student))
	}
}
