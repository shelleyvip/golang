package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {

	//第一步 首先输入一个范围（指定爬取起始页到终止页 一个1-多少的范围）
	var start, end int
	fmt.Printf("请输入起始页( >= 1) :")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页( >= 起始页) :")
	fmt.Scan(&end)

	Dowork1(start, end)
}
func Dowork1(start, end int)() {
	fmt.Printf("正在爬取 %d 到 %d 的页面\n", start, end)

	//循环读取每一页数据
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E4%B8%89%E5%92%8C%E5%A4%A7%E7%A5%9E&ie=utf-8&pn=" +
			strconv.Itoa((i-1)*50) //// 由于它从一开始输入，所以要-1 *50 
		//2）//把爬取的内容，返回到result 用result去接收
		result, err := HttpGet6(url) //
		if err != nil {
			fmt.Println("HttpGet err", err)
			continue
		}
		//将读到的整网页数据  保存成一个文件
		f,err := os.Create("第"+strconv.Itoa(i)+"页"+".html")
		if err != nil{
			fmt.Println("os.Create err",err)
			continue
		}
		f.WriteString(result)
		f.Close()  //保存好一个文件，关闭一个文件
	}
}
func HttpGet6(url string)(result string,err error)  {
	resp,err1 := http.Get(url)
	if err1 != nil{
		err = err1 //将封装函数的内部错误，传出给调用者
		return
	}
	defer resp.Body.Close()

	//循环读取页年数据，传出给调用者
	buf :=make([]byte,4096)
	for{
		n,err2 := resp.Body.Read(buf)
		if n == 0{
			fmt.Println("读取网页完成")
			break
		}
		if err2 != nil && err2 == io.EOF{
		err = err2
		return
		}
		//累加每一次循环 读到的 buf数据，存入result 一次性返回。
		result += string(buf[:n])
	}
   return //把它返回回去
}