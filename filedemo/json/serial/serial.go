package serial

import (
	"encoding/json"
	"fmt"
)

//先声明定义一个结构体
type Monster struct {
	Name string   `json:"monster_name"`   //反序列化
	Age int `json:"monster_age"`
	Birthday string
	sal float64
	skill string
}

func TestStruct()  {
	//演示
	monster := Monster{
		Name : "牛魔王",
		Age : 88,
		Birthday: "1988-2-2",
		sal:8000.00,
		skill:"吹牛皮",
	}
	//将monster序列化
	data,err := json.Marshal(&monster)
	if err != nil{
		fmt.Printf("序列号错误%v\n",err)
	}
	//输出序列化的结果
	fmt.Printf("monster序列化后=%v\n",string(data))
}

//将map进行序列化
func testMap()  {

	//定义一个map是string类型 值是任意类型
	//他的key 是string 值是interface 任意类型
	var a map[string]interface{}

	//使用map先make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "火云洞"

	////将testMap序列化
	data,err := json.Marshal(a)
	if err != nil{
		fmt.Printf("序列化错误=err%v\n",err)
	}
	fmt.Printf("testMap序列化后=%v\n",string(data))
}

//演示对切片进行序列化
func testSlice()  {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	//使用map前先make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1 ["age"] =9
	m1["address"] = "北京"

	slice = append(slice,m1)

	//
	var m2 map[string]interface{}
	//使用map前先make
	m2 = make(map[string]interface{})
	m2["name"] = "jack2222"
	m2 ["age"] =   88
	m2["address"] = "上海"

	slice = append(slice,m2)

		}
//对基本数据类型序列化
func testFloat64()  {
	var num1 float64= 2814.22
	//对num1进行序列化
	data,err :=json.Marshal(num1)
	if err != nil{
		fmt.Printf("序列化错误err=%v\n",err)
	}
	fmt.Printf("testFloat64 序列化后=%v\n",string(data))
}



func main() {
	//演示将结构体，map，切片进行序列化
	TestStruct()
	testMap()
	testSlice()    //演示对切片的序列化
	testFloat64()  //演示对基本数据类型序列化
}
