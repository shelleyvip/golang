package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name string
	Age int
	Skill string
}


//给monster绑定一个Store的方法，可以将monster变量（对象），序列化后保存在文件中

func (this *Monster)Store() bool {
	//1:先序列化
	data,err := json.Marshal(this)
	if err != nil{
		fmt.Println("marshal err=",err)
		return false
	}
	//2:保存在文件

	filePath := "/Users/golang/desktop/shelley.exe"
	err = ioutil.WriteFile(filePath,data,0666)
	if err!= nil{
		fmt.Println("writer	file err=",err)
		return false
	}
	return true
}


//给Monster绑定一个方法ReStore,可以将一个序列化的Monster,从文件中读取
//并反序列化为Monster对象，检查反序列化，名字正确

func (this *Monster)ReStore()bool  {
	//1:先从文件中，读取序列化的字符串
	filePath := "/Users/golang/desktop/shelley.exe"
	data,err := ioutil.ReadFile(filePath)
	if err != nil{
		fmt.Println("readFile err=",err)
		return false
	}
	//2:使用读取到data[]byte,对反序列化
	err = json.Unmarshal(data,this)//第一个是数据，第二个是恢复到哪个地方去
	if err != nil{
		fmt.Println("Unmarshal err=",err)
		return false
	}
	return true
}
