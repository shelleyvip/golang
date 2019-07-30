package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//首次定一个结构体 用于保存统计的结果
type CharCount struct {
    ChCount int  //记录英文的个数
    NumCount int //记录数字的个数
    SpaceCount int  //记录空格的个数吧
    OtherCount int //记录其它字符

}

func main() {
	//思路：打开一个文件，创建一个reader
	//每读取一行，就去统计改行有多少个 英文 数字 空格 个其它字符
	//然后将结果保存到一个结构体中
	fileName := "/Users/golang/desktop/aaa.txt"
	file,err :=os.Open(fileName)
	if err != nil{
		fmt.Printf("open file err=%v\n",err)
		return
	}
	defer file.Close()

	//定一个结构体CharCount的实例
	var count CharCount

	//创建一个reader
	reader := bufio.NewReader(file) //把文件句柄放进去

	//开始循环的读取fileName文件内容

	for{
		str,err := reader.ReadString('\n')
		if err == io.EOF{ //如果读到文件的末尾 就不在读了
		break
		}
		//这个str就是我们读到的字符串。遍历一下
		//遍历str 进行统计
		for _,v := range str{

			switch {
			case v >= 'a' && v <= 'z': //以内字母就是a-z
                     fallthrough //穿透 继续执行
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++

			}
		}
	}
    //输出统计结果看看属否正确
    fmt.Printf("字符的个数为:%v 数字的个数为:%v 空格的个数为:%v 其它字符个数:%v",
    	count.ChCount,count.NumCount,count.SpaceCount,count.OtherCount)
}
