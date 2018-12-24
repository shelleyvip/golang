package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//指定爬取起，终止页
	var start, end int
	fmt.Printf("请输入起始页(>= 1):")//提示用户输入
	fmt.Scan(&start)
	fmt.Printf("请输入终止页（>= 终止页）:")
	fmt.Scan(&end)

	Dowork(start, end) //工作函数
}
//爬取页面操作
func Dowork(start, end int) {
	fmt.Printf("正在爬取第%d页面%d \n", start, end)
	//循环爬取每一页
	for i := start; i <= end; i++ {
		//定义一个函数爬取主页面
		 Spiderpape(i)
	}
}
func Spiderpape(i int) {
	//明确爬取的Url
	//https://www.pengfu.com/xiaohua_1.html
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬取第%d页面%s \n", i, url)

	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err", err)
		return
	}
	//取<h1 class="dp-b"><a href="一个段子URL连接
	//解析表达式
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if re == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}

	//取关键信息   -1带表取所有内容
	joyUrls := re.FindAllStringSubmatch(result, -1)
	//fmt.Println("joyUrls =",joyUrls)
	fileTitle := make([]string, 0)
	fileContent := make([]string, 0)
	//取网址
	//第一个返回的是下标，第二个返回内容
	for _, data := range joyUrls {
		//fmt.Println("url=",data[1])
		//开始爬取每一个笑话，每一个段子
		title, content, err := SpiderOneJoy(data[1])
		if err != nil {
			continue
			fmt.Println("SpiderOneJoy err", err)
		}
		//fmt.Printf("title=#%v#",title)
		//fmt.Printf("content=#%v#",content)
		fileTitle = append(fileTitle, title) //追加内容
		fileContent = append(fileContent, content)
	}
	//fmt.Println("fileTitle",fileTitle)
	//fmt.Println("fileContent ",fileContent)
	StoteJoyToFile(i, fileTitle, fileContent)

}
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url) //发送get请求
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024*4)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result += string(buf[:n]) //累加读取的内容
	}
	return
}

////开始爬取每一个笑话，每一个段子
//		title,content,err := SpiderOneJoy(url)
func SpiderOneJoy(url string) (title, content string, err error) {
	//开始爬取内容
	result, err1 := HttpGet(url)
	if err1 != nil {
		err = err1
		return
	}
	//取关键信息
	//取标题
	re1 := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if re1 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile err1")
		return
	}
	//取内容
	tmpTitle := re1.FindAllStringSubmatch(result, 1) //最后一个参数为1，只过滤第一个
	for _, data := range tmpTitle {
		title = data[1]
		title = strings.Replace(title, "\t ", "", -1)
		break
	}
	//取内容
	re2 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	if re2 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile err2")
		return
	}
	//取内容
	tmpContent := re2.FindAllStringSubmatch(result, -1)
	for _, data := range tmpContent {
		content = data[1]
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "\n", "", -1)
		content = strings.Replace(content, "\r", "", -1)
		content = strings.Replace(content, "<br />", "", -1)

		break
	}
	return
}
func StoteJoyToFile(i int, fileTitle, fileContent []string) {
	f, err := os.Create(strconv.Itoa(i) + ".txt")
	if err != nil {
		fmt.Println("os.Create err", err)
		return
	}
	defer f.Close()
	//写内容
	n := len(fileTitle)
	for i := 0; i < n; i++ {
		f.WriteString(fileTitle[i] + "\n")   //写标题
		f.WriteString(fileContent[i] + "\n") //写内容
		f.WriteString("\n=================\n")
	}

}
