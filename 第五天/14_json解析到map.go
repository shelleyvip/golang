package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	jsonBuf :=`
      {
		"Company": "itcast",
		"Subjects":	[
          "Go",
         "C++",
         "Python",
         "Test"
        ],
      "IsOk":true,
      "Price":666.666
}`
	//创建一个map
	m := make(map[string]interface{},4)
	err := json.Unmarshal([]byte(jsonBuf),&m) //字节码
	if err != nil{
		fmt.Println("err=",err)
		//log.Fatal(err)
		return
	}
		fmt.Printf("m=%+v\n",m)


	var str string
	//类型断言，值，它是value类型
	for key,value := range m{
		switch data := value.(type){
		case string:
			str = data
			fmt.Printf("map[%s]的值类型为string,value=%s\n",key,str)
		case bool:
			fmt.Printf("map[%s]的值类型为bool,value=%v\n",key,data)
		case float64:
			fmt.Printf("map[%s]的值类型为float64,value=%v\n",key,value)
		case interface{}:
			fmt.Printf("mp[%s]的值类型为intarfece,value=%v\n",key,value)

		}
	}



}
