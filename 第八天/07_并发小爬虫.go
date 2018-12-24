package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	//第一步 首先输入一个范围（1-多少的范围）
	var start, end int
	fmt.Printf("请输入起始页（>= 1）:")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页(>= 起始页):")
	fmt.Scan(&end)
	DoworkA(start, end)
}

func DoworkA(start, end int) {
	fmt.Printf("正在爬取%d到%d的页面", start, end)
	//1)明确目标（要知道你准备在哪个范围或者哪个网站去搜索
	//"https://tieba.baidu.com/f?kw=%E8%B6%85%E8%83%BD%E5%8A%9B&ie=utf-8&pn=0 下一页加50
	page := make(chan int)
	for i := start; i <= end; i++ {
		go SpidePepa(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}
}
func HttpGetA(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		fmt.Println("http.Get err", err1)
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024*4)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
			//fmt.Println("resp.Body.Read err", err)
		}
		if err2 != nil && err2 == io.EOF {
			return

		}
		result += string(buf[:n])

	}

	return

}
//爬一个网页
func SpidePepa(i int, page chan<- int) {
	url := "https://tieba.baidu.com/f?kw=%E8%B6%85%E8%83%BD%E5%8A%9B&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	fmt.Printf("正在爬%d页网页%s\n", i, url)


	// 2）爬 将网站的所有的内容全部爬下来
	result, err := HttpGetA(url)
	if err != nil {
		fmt.Println("HttpGet err", err)
		return
	}
	//把内容放在文件夹
	FileName := strconv.Itoa(i) + ".html"
	f, err1 := os.Create(FileName)
	if err1 != nil {
		fmt.Println("os.Create")
		return
	}
	f.WriteString(result) //写内容

	f.Close()  //关闭
	page <- i

}