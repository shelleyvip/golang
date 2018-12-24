package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

//爬虫概念：
    //访问web服务器，获取指定信息的一段程序

//工作流程：
  //1：明确目标URl
  //2：发送请求 获取应答数据包
   //3：保存过滤数据，提取有用的信息
   //4 :使用 分析得到的数据

//百度贴吧爬虫的实现：操作流程如下
//1：提示用户指定起始，终止页。创建working()函数
//2：使用start，end循环爬取每一页数据
//3：获取每一页的URL---下一页=前一页+50
//4：封装 实现HttpGet（）函数，爬取一个网页的数据内容，通过result 返回
// http.Get/resp.Body.Close/buf:=make(1024*4)/for{resp.Body.Read(buf)/result += string(buf[:n])return
//5:创建 .html文件。使用循环因子i命名
//6:result写入 文件WriteString(result).f.close 不推荐使用 defer


//并发版百度爬虫:
//1:封装爬取一个网页的内容到SpiderPage(index)函数中
//2:在working 函数 for循环启动go 程 循环调用SpiderPage（）————>n个待爬取页面，对应n个go程
//3:为防止主go程提前结束，引入channel实现同步。SpiderPage（index，channel）
//4:在SpiderPage结尾处（一个页面爬去完成），向channel中写入内容。channel <- index
//5:在working函数添加一个for循环，从channel不断读取各个子go程写入的数据，n个子go程---写n次channel---读n次channel

//爬去网页内
func HttpGet2(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1 //将封装函数内部的错误，传出给调用者，（调用者也就是HttpGet）
		return
	}
	defer resp.Body.Close()

	//循环读取数据，传出给调用者
	buf := make([]byte, 1024*4)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 { // 读取结束或者出问题
			break
		}
		if err2 != nil && err2 == io.EOF{
			return
		}
		//累加每一次循环 读到的 buf数据，存入result 一次性返回。
		result += string(buf[:n])
	}
	return
}

//爬取单个页面的函数
func Spiderpage(i int,page chan int) {
	url := "https://tieba.baidu.com/f?kw=%E4%B8%89%E5%92%8C%E5%A4%A7%E7%A5%9E&ie=utf-8&pn=" +
		strconv.Itoa((i-1)*50) //// 由于它从一开始输入，所以要-1 *50
	//2）//把爬取的内容，返回到result 用result去接收
	result, err := HttpGet2(url) //
	if err != nil {
		fmt.Println("HttpGet err", err)
		return
	}
	//将读到的整网页数据  保存成一个文件
	f,err := os.Create("第"+strconv.Itoa(i)+"页"+".html")
	if err != nil{
		fmt.Println("os.Create err",err)
		return
	}
	f.WriteString(result)
	f.Close()  //保存好一个文件，关闭一个文件
	page <- i //与主fgo程完成同步
}


//爬取页面操作
func working(start, end int) {
	fmt.Printf("正在爬取 %d 到 %d 的页面\n", start, end)

	page := make(chan int)
	//循环爬取每一页数据
	for i:=start;i<=end;i++{
		go Spiderpage(i,page)
	}
	for i:=start;i<=end;i++{
		fmt.Printf("第%d个页面爬取完成\n",<-page)
	}


}

func main() {

	//第一步 首先输入一个范围（指定爬取起始页到终止页 一个1-多少的范围）
	var start, end int
	fmt.Printf("请输入起始页( >= 1) :")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页( >= 起始页) :")
	fmt.Scan(&end)

	working(start, end)
}