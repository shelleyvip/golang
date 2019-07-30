package main

import (
	"encoding/json"
	"fmt"
)
//定义一个结构体
type Monster struct {
	Name string
	Age int
	Birthday string

}

//演示    //将json字符串反序列化成struct
func unmarshalStruct()  {
	//说明：这个str在项目开发中 是通过网络传输获取到的。。或者是读取文件 获取到的
       str := "{\"Name\":\"牛魔王\",\"Age\":88,\"Birthday\":\"1988-2-2\"}"

       //定义一个Monster实例
       var monster Monster

       //这里面会接收两个参数一个是str里面的字串所对应的byte切片，另外一个是结构体对应的字符串
       err := json.Unmarshal([]byte(str),&monster)  //另外一个是反序列化过后交给的是哪个变量
       if err !=nil{
       	fmt.Printf("unmarshal err=%v\n",err)
	   }
       fmt.Printf("反序列化后 moster=%v",monster,monster.Birthday)

}


func main() {
	unmarshalStruct()
}
