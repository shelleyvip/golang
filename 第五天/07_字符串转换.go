package main

import (
	"fmt"
	"strconv"
)

func main(){
	//转换到字符串后 追加到字节数组
	slice := make([]byte,0,1234)
	slice = strconv.AppendBool(slice,true)
	//第二个为要追加的数，第三个为指定 以10进制方式追加
	slice = strconv.AppendInt(slice,888,10)
	slice = strconv.AppendQuote(slice,"abcgohello")

	fmt.Println("slice = ",string(slice))//转换为string后在进行打印

	//字符串转其他类型
	var flag bool
	var err error
	flag,err = strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag=",flag)
	}else{
		fmt.Println("err=",err)
	}
	//其它类型转换为字符串
	var str string //字符串转bool类型 bool 只有两个值 false和true
	str = strconv.FormatBool(false)
	fmt.Println("str=",str)

	//'f'指打印格式，以小数方式，—1指小数点位数（紧缩模式），64 以float64处理
	abc := strconv.FormatFloat(3.14,'f',-1,64)
	fmt.Println("abc=",abc)


	//整型转化为字符串，常用
	var tmp string
	tmp = strconv.Itoa(999)
	fmt.Println("tmp=",tmp)


	//把字符串转换为整型
	a,_ := strconv.Atoi("567")
	fmt.Println("a=",a)
}