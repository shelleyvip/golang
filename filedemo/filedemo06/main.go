package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//将某某文件的内容导入到 某某文件中

	//首先将某某盘 内容读取到 某某盘
	//将读取到的内容读取到内存
	fili1Path := "/Users/golang/work/go_path/src/golang/filedemo/shelley.txt"
	file2Path :=  "/Users/golang/work/go_path/src/golang/xxx.exe"

	//先读出来
	data,err := ioutil.ReadFile(fili1Path)
	if err != nil{
		//说明文件有错
		fmt.Printf("read file err =%v",err)
		return
	}
	err = ioutil.WriteFile(file2Path,data,666)
	if err != nil{
		fmt.Printf("writer file err=%v",err)

	}


}
