package main
import (
"encoding/json"  //结构体生成 json 的包
"fmt"
)

//成员首字母必须大写
type IT struct {
Company string
Subjects []string
IsOk bool
Price float64
}

//定义一个结构体变量
func main(){
s1 := IT{"shelley",[]string{"Go","C++","Python","Test"},true,666.666}
// 根据内容生成一个json文本
//buf,err := json.Marshal(s1)  //这是第一种看上去不清晰 下面是它打印的结果
//buf= {"Company":"shelley","Subjects":["Go","C++","Python","Test"],"IsOk":true,"Price":666.666}

//这个就是编码了， 根据内容生成一个json文本
buf,err := json.MarshalIndent(s1,""," ")//这两个一个是空"" 一个是" "缩进。这就是格式化编码：重要！
if err != nil{       //indent 就是缩进的意思
fmt.Println("err=",err)
}else {
fmt.Println("buf=",string(buf))
}
//这里就是格式打印出来的结果
/*buf= {
	"Company": "shelley",
		"Subjects": [
	"Go",
		"C++",
		"Python",
		"Test"
	],
	"IsOk": true,
		"Price": 666.666
}

*/}