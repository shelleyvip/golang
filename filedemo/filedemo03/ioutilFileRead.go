package main

import (
	"fmt"
	"io/ioutil"    //这种一次性读取的 只适用于文件小的内容
)

func main() {
	//使用ioutil.ReadFile一次性将文件读取到位
	file := "/Users/golang/desktop/复习秘籍.html"
	content,err := ioutil.ReadFile(file)
	if err != nil{
		fmt.Printf("read file err=%v",err)
	}
	//fmt.Printf("%v",content)   //这里要转成字符串输出
	fmt.Printf("%v",string(content))
	//把读取的内容显示在终端  //[]byte
   //我们没有显示的open file 所以也不需要显示的close file
   //因为 open和close 北封装到read file函数内部了
}
