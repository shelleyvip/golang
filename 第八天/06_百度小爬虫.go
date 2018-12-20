package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

//爬去网页内
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	//读取网页body 内容
	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 { //读取结束或者出问题
			fmt.Println("resp.Body.Read err", err)
			break
		}
		result += string(buf[:n])

	}
	return
}



func Dowork(start, end int) {
	fmt.Printf("正在爬取 %d 到 %d 的页面\n", start, end)

	//1)明确目标（要知道你准备在哪个范围或者哪个网站去搜索）
	for i := start; i < end; i++ { ///第二步✌️然后根据这个范围来个循环来生成网页 网址，
		url := "https://tieba.baidu.com/f?kw=%E4%B8%89%E5%92%8C%E5%A4%A7%E7%A5%9E&ie=utf-8&pn=" +
			strconv.Itoa((i-1)*50) //// 由于它从一开始输入，所以要-1 *50
		fmt.Println("url =", url)

		//2）//爬(将所有网站的内容全部爬去下来)
		result, err := HttpGet(url) //  第三步 ：：把网址放进去 //爬完之后反馈结果result
		if err != nil {
			fmt.Println("HttpGet err", err)
			continue
		}

		//把内容写入到文件
		FileName := strconv.Itoa(i) + ".html"
		f, err1 := os.Create(FileName)   //第四步 ：：往这个文件写内容
		if err1 != nil {
			fmt.Println("os.Create err1", err1)
			continue
		}

		f.WriteString(result) //写内容
		f.Close()//关闭文件
	}

}

func main() {

	//第一步 首先输入一个范围（1-多少的范围）
	var start, end int
	fmt.Printf("请输入起始页( >= 1) :")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页( >= 起始页) :")
	fmt.Scan(&end)

	Dowork(start, end)
}