package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//创建一个map
	m := make(map[string]interface{},4)
	m["company"] = "itcast"
	m["subjects"] = []string{"GO","C++","Python","Tst"}
	m["isok"] = true
	m["Test"] = 3.14

	result,err := json.MarshalIndent(m,""," ")
	if err != nil{
		fmt.Println("err=",err)
		return
	}
	fmt.Println("result=",string(result))
}
