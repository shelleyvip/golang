package main

import (
	"encoding/json"   //map 生成 json 的包
	"fmt"
)

func main(){
	//创建一个map
	m := make(map[string]interface{},4)
	m["company"]= "itcast"
	m["subjects"]= []string{"Go","C++","Python","Test"}
	m["isOk"] = true
	m["price"] = 666.666
///变成json                          ""，" " 这就是格式化编码：重要！
	result,err := json.MarshalIndent(m,""," ")//一个是空"" 第二是" "缩进。
	if err != nil{
		fmt.Println("err=",err)
		return
	}else {
		fmt.Println("result=",string(result))
	}
}
