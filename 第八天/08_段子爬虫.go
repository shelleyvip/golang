package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	var start,end int
	fmt.Printf("请输入起始页(>= 1):")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页（>= 终止页）:")
	fmt.Scan(&end)

	Dowork(start,end)  //工作函数
}
func Dowork(start,end int)  {
	fmt.Printf("正在爬取第%d页面%d \n",start,end)
	  for i :=start;i <= end;i++{
	  	//定义一个函数爬取主页面
	  	Spiderpape(i)
	}
}
func Spiderpape(i int)  {
	//明确爬取的Url
	//https://www.pengfu.com/xiaohua_1.html
	url := "https://www.pengfu.com/xiaohua_"+strconv.Itoa(i)+".html"
	fmt.Printf("正在爬取第%d页面%s \n",i,url)

	result,err := HttpGet(url)
	if err != nil{
		fmt.Println("HttpGet err",err)
		return
	}
	//取<h1 class="dp-b"><a href="一个段子URL连接
	//解释表达式
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if re == nil{
		return
		fmt.Println("regexp.MustCompile err")
	}

	//取关键信息
	joyUrls := re.FindAllStringSubmatch(result,-1)
	//fmt.Println("joyUrls =",joyUrls)
	//取网址
	//第一个返回的是下标，第二个返回内容
	for _,data := range joyUrls{
		//fmt.Println("url=",data[1])
		//开始爬取每一个笑话，每一个段子
		title,content,err := SpideeOneJoy(url)
	}

	}
func HttpGet(url string)(result string,err error)  {
	resp,err1 := http.Get(url)//发送get请求
	if err1 != nil{
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte,1024*4)
	for{
		n,_ := resp.Body.Read(buf)
		if n == 0{
			break
		}
		result += string(buf[:n]) //累加读取的内容
	}
return
}